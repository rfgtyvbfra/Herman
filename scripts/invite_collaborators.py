#!/usr/bin/env python3
import csv
import json
import sys
import urllib.request
import urllib.error
from pathlib import Path
from typing import Optional, Dict

ROOT = Path(__file__).resolve().parents[1]
CSV_PATH = ROOT / 'github_accounts.csv'
OWNER = 'rfgtyvbfra'
REPO = 'Herman'


def load_accounts():
    with open(CSV_PATH, newline='') as f:
        return list(csv.DictReader(f))


def owner_token(rows):
    for r in rows:
        if r['username'] == OWNER:
            return r['token']
    raise SystemExit('owner token not found')


def api_request(method: str, url: str, token: str, data: Optional[Dict] = None):
    req = urllib.request.Request(url, method=method)
    req.add_header('Accept', 'application/vnd.github+json')
    req.add_header('Authorization', f'token {token}')
    if data is not None:
        body = json.dumps(data).encode('utf-8')
        req.add_header('Content-Type', 'application/json')
    else:
        body = None
    try:
        with urllib.request.urlopen(req, body) as resp:
            payload = resp.read().decode('utf-8')
            return resp.status, payload
    except urllib.error.HTTPError as e:
        payload = e.read().decode('utf-8')
        return e.code, payload


def add_collaborator(token: str, username: str):
    url = f'https://api.github.com/repos/{OWNER}/{REPO}/collaborators/{username}'
    status, payload = api_request('PUT', url, token, data={"permission": "push"})
    print(f'add_collaborator {username}: HTTP {status}')


def main():
    rows = load_accounts()
    token = owner_token(rows)
    for r in rows:
        u = r['username']
        if u == OWNER:
            continue
        add_collaborator(token, u)


if __name__ == '__main__':
    main()
