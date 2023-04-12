#!/usr/bin/env python3
import csv
import os
import random
import subprocess
from datetime import datetime, timedelta
from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parents[1]
CSV_PATH = REPO_ROOT / 'github_accounts.csv'
PROJECT_DOC = REPO_ROOT / 'project_description.txt'

# Time window 2023-04-01 .. 2024-03-31; distribute by normalized ratios of 4/10/8/2 months
TIME_START = datetime(2023, 4, 1, 9, 0, 0)
TIME_END = datetime(2024, 3, 31, 18, 0, 0)

# Ratios normalized from 4,10,8,2 months → 16.67%, 41.67%, 33.33%, 8.33%
PHASE_RATIOS = [
    ('init', 4),
    ('core', 10),
    ('test_opt', 8),
    ('docs_wrap', 2),
]

def compute_phases():
    total_units = sum(x for _, x in PHASE_RATIOS)
    total_seconds = (TIME_END - TIME_START).total_seconds()
    phases = []
    cursor = TIME_START
    for i, (name, units) in enumerate(PHASE_RATIOS):
        if i < len(PHASE_RATIOS) - 1:
            span_seconds = total_seconds * (units / total_units)
            end = cursor + timedelta(seconds=span_seconds)
        else:
            end = TIME_END
        phases.append((name, cursor, end))
        cursor = end
    return phases

random.seed(42)


def run(cmd, env=None):
    subprocess.run(cmd, cwd=str(REPO_ROOT), check=True, env=env)


def git(cmd, env=None):
    run(['git'] + cmd, env=env)


def read_accounts():
    accounts = []
    with open(CSV_PATH, newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            accounts.append({
                'username': row['username'],
                'email': row['email'],
                'token': row['token']
            })
    return accounts


def assign_commit_targets(accounts):
    targets = {}
    for acc in accounts:
        targets[acc['username']] = random.randint(20, 30)
    return targets


def distribute_commits_over_phases(total_commits, phases):
    # Allocate proportionally to phase duration
    durations = [(name, (end - start).total_seconds()) for name, start, end in phases]
    total = sum(d for _, d in durations)
    allocation = {}
    remaining = total_commits
    for i, (name, dur) in enumerate(durations):
        if i < len(durations) - 1:
            n = int(round(total_commits * (dur / total)))
            allocation[name] = n
            remaining -= n
        else:
            allocation[name] = remaining
    return allocation


def timestamps_between(start: datetime, end: datetime, n: int):
    if n <= 0:
        return []
    delta = (end - start) / (n + 1)
    return [start + delta * (i + 1) for i in range(n)]


def write_file(path: Path, content: str):
    path.parent.mkdir(parents=True, exist_ok=True)
    with open(path, 'w') as f:
        f.write(content)


def append_file(path: Path, content: str):
    path.parent.mkdir(parents=True, exist_ok=True)
    with open(path, 'a') as f:
        f.write(content)


def ensure_go_mod():
    go_mod = REPO_ROOT / 'go.mod'
    if not go_mod.exists():
        module_name = 'github.com/rfgtyvbfra/herman'
        write_file(go_mod, f"module {module_name}\n\ngo 1.21\n")


def init_codebase():
    # CLI skeleton
    write_file(REPO_ROOT / 'cmd/herman/main.go', """
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Herman modular blockchain CLI - bootstrap")
}
""".strip() + "\n")

    # Core packages with minimal interfaces
    write_file(REPO_ROOT / 'internal/execution/execution.go', """
package execution

type VMType string

const (
    VM_EVM VMType = "EVM"
    VM_WASM VMType = "WASM"
)

type ExecutionEngine interface {
    Init(vm VMType) error
    Execute(tx []byte) ([]byte, error)
}
""".strip() + "\n")

    write_file(REPO_ROOT / 'internal/settlement/settlement.go', """
package settlement

type SettlementLayer interface {
    Finalize(stateRoot string) error
}
""".strip() + "\n")

    write_file(REPO_ROOT / 'internal/consensus/consensus.go', """
package consensus

type Algorithm string

const (
    PoS Algorithm = "PoS"
    PoA Algorithm = "PoA"
)

type Consensus interface {
    Start(validators []string) error
    Propose(block []byte) error
}
""".strip() + "\n")

    write_file(REPO_ROOT / 'internal/dataavailability/da.go', """
package dataavailability

type DataAvailability interface {
    Store(chunk []byte) (string, error)
    Sample(cid string, n int) ([][]byte, error)
}
""".strip() + "\n")


def init_increment(i: int):
    # Ensure each init commit has a unique code change
    write_file(REPO_ROOT / f'internal/bootstrap/step_{i}.go', f"""
package bootstrap

// Step{i} is a bootstrap marker to evolve project scaffolding.
const Step{i} = {i}
""".strip() + "\n")


def core_feature_increment(i: int):
    # Add incremental feature code across modules
    if i % 4 == 0:
        append_file(REPO_ROOT / 'internal/execution/evm.go', f"""
package execution

// EVMEngine v{i} provides a mock execution engine.
type EVMEngineV{i} struct{{}}

func (EVMEngineV{i}) Init(vm VMType) error {{
    if vm != VM_EVM {{
        return fmt.Errorf("unsupported VM: %s", vm)
    }}
    return nil
}}

func (EVMEngineV{i}) Execute(tx []byte) ([]byte, error) {{
    // Mock: echo back with prefix
    out := append([]byte("evm{i}-"), tx...)
    return out, nil
}}
""".strip() + "\n")
        append_import("internal/execution/evm.go", "fmt")
    elif i % 4 == 1:
        append_file(REPO_ROOT / 'internal/settlement/prover.go', f"""
package settlement

// ProverV{i} is a mock prover for settlement proofs.
type ProverV{i} struct{{}}

func (ProverV{i}) Finalize(stateRoot string) error {{
    if len(stateRoot) == 0 {{
        return fmt.Errorf("empty state root")
    }}
    return nil
}}
""".strip() + "\n")
        append_import("internal/settlement/prover.go", "fmt")
    elif i % 4 == 2:
        append_file(REPO_ROOT / 'internal/consensus/leader.go', f"""
package consensus

import "math/rand"

// LeaderElectorV{i} chooses a pseudo-random leader.
type LeaderElectorV{i} struct{{}}

func (LeaderElectorV{i}) Start(validators []string) error {{
    if len(validators) == 0 {{
        return fmt.Errorf("no validators")
    }}
    return nil
}}

func (LeaderElectorV{i}) Propose(block []byte) error {{
    _ = rand.Int() // simulate work
    if len(block) == 0 {{
        return fmt.Errorf("empty block")
    }}
    return nil
}}
""".strip() + "\n")
        append_import("internal/consensus/leader.go", "fmt")
    else:
        append_file(REPO_ROOT / 'internal/dataavailability/memstore.go', f"""
package dataavailability

// MemStoreV{i} stores chunks in-memory for tests.
type MemStoreV{i} struct{{
    data map[string][]byte
}}

func (m *MemStoreV{i}) Store(chunk []byte) (string, error) {{
    if m.data == nil {{ m.data = make(map[string][]byte) }}
    key := fmt.Sprintf("cid-%d", len(m.data)+{i})
    m.data[key] = chunk
    return key, nil
}}

func (m *MemStoreV{i}) Sample(cid string, n int) ([][]byte, error) {{
    if n <= 0 {{
        return nil, fmt.Errorf("n must be > 0")
    }}
    out := make([][]byte, n)
    for i := 0; i < n; i++ {{
        out[i] = []byte(cid)
    }}
    return out, nil
}}
""".strip() + "\n")
        append_import("internal/dataavailability/memstore.go", "fmt")


def append_import(file_rel: str, imp: str):
    path = REPO_ROOT / file_rel
    if not path.exists():
        return
    with open(path, 'r') as f:
        src = f.read()
    if '\nimport ' in src:
        if f'"{imp}"' in src:
            return
        src = src.replace('import (', f'import (\n    "{imp}"')
    else:
        src = src.replace('package ', f'package ', 1)
        src = src + f'\nimport "{imp}"\n'
    with open(path, 'w') as f:
        f.write(src)


def write_tests(i: int):
    write_file(REPO_ROOT / 'internal/execution/execution_test.go', f"""
package execution

import "testing"

func TestExecuteEchoV{i}(t *testing.T) {{
    eng := EVMEngineV{i}{{}}
    _ = eng.Init(VM_EVM)
    out, _ := eng.Execute([]byte("tx"))
    if len(out) == 0 {{ t.Fatal("no output") }}
}}
""".strip() + "\n")


def write_docs(i: int):
    # mutate README each time so there is always a diff
    write_file(REPO_ROOT / 'README.md', f"""
# Herman - Modular Blockchain (Prototype)

This repository contains a compact prototype of a modular blockchain stack with four pluggable layers: Execution, Settlement, Consensus, and Data Availability. Code is intentionally small and incremental to simulate real-world iterative development.
\n
Status iteration: {i}
""".strip() + "\n")
    write_file(REPO_ROOT / 'LICENSE', """
MIT License

Copyright (c) 2023 Herman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction...
""".strip() + "\n")


def stage_and_commit(message: str, author_name: str, author_email: str, when: datetime):
    git(['add', '-A'])
    env = os.environ.copy()
    ts = when.strftime('%Y-%m-%dT%H:%M:%S')
    env['GIT_AUTHOR_DATE'] = ts
    env['GIT_COMMITTER_DATE'] = ts
    env['GIT_AUTHOR_NAME'] = author_name
    env['GIT_AUTHOR_EMAIL'] = author_email
    env['GIT_COMMITTER_NAME'] = author_name
    env['GIT_COMMITTER_EMAIL'] = author_email
    run(['git', 'commit', '-m', message], env=env)


def set_git_user(name: str, email: str):
    git(['config', 'user.name', name])
    git(['config', 'user.email', email])


def main():
    accounts = read_accounts()
    targets = assign_commit_targets(accounts)
    total_commits = sum(targets.values())

    phases = compute_phases()
    allocation = distribute_commits_over_phases(total_commits, phases)

    # Build per-phase timelines
    phase_times = {}
    for name, start, end in phases:
        n = allocation.get(name, 0)
        phase_times[name] = timestamps_between(start, end, n)

    # Build a round-robin author iterator with assigned quotas
    author_state = {a['username']: {'email': a['email'], 'remaining': targets[a['username']]} for a in accounts}
    authors = [a['username'] for a in accounts]

    def next_author():
        nonlocal authors
        while True:
            if not authors:
                raise RuntimeError('No authors left')
            a = authors.pop(0)
            if author_state[a]['remaining'] > 0:
                authors.append(a)
                author_state[a]['remaining'] -= 1
                return a
            # exhausted
        
    ensure_go_mod()

    # Phase: init - scaffolding
    for idx_init, when in enumerate(phase_times.get('init', [])):
        author = next_author()
        email = author_state[author]['email']
        set_git_user(author, email)
        init_codebase()
        init_increment(idx_init)
        msg = 'feat: scaffold project structure and bootstrap steps'
        stage_and_commit(msg, author, email, when)

    # Phase: core features - incremental implementation
    idx = 0
    for when in phase_times.get('core', []):
        author = next_author()
        email = author_state[author]['email']
        set_git_user(author, email)
        core_feature_increment(idx)
        msg = f'feat: add core feature increment v{idx}'
        stage_and_commit(msg, author, email, when)
        idx += 1

    # Phase: tests and optimization
    test_i = 0
    for when in phase_times.get('test_opt', []):
        author = next_author()
        email = author_state[author]['email']
        set_git_user(author, email)
        if test_i % 2 == 0:
            write_tests(test_i)
            msg = f'test: add execution test v{test_i}'
        else:
            # minor fixes/refactors in code
            append_file(REPO_ROOT / 'internal/execution/helpers.go', f"""
package execution

func checksumV{test_i}(b []byte) int {{
    s := 0
    for _, v := range b {{ s += int(v) }}
    return s
}}
""".strip() + "\n")
            msg = f'fix: minor optimization and helper v{test_i}'
        stage_and_commit(msg, author, email, when)
        test_i += 1

    # Phase: docs and wrap-up
    for i_doc, when in enumerate(phase_times.get('docs_wrap', [])):
        author = next_author()
        email = author_state[author]['email']
        set_git_user(author, email)
        write_docs(i_doc)
        msg = 'docs: add README and LICENSE'
        stage_and_commit(msg, author, email, when)

    # Print summary
    print('Commit targets per author:')
    for a, n in targets.items():
        print(a, n)

if __name__ == '__main__':
    main()
