# packetfabric_streaming_events.py
import argparse
import datetime
import json
import os
import threading
import requests
from sseclient import SSEClient


def logout(token, base_url):
    headers = {"Authorization": f"Bearer {token}"}
    logout_url = f"{base_url}/v2/auth/logout"
    requests.get(logout_url, headers=headers)
    print("Logged out")

def consume_client(token, base_url, subscription_uuid, output_file, duration_seconds):
    headers = {"Authorization": f"Bearer {token}", "Accept": "text/event-stream"}
    url = f"{base_url}/v2/events/{subscription_uuid}"
    r = SSEClient(url, headers=headers)
    events = []

    start_time = datetime.datetime.utcnow()

    for event in r:
        if event.data.strip():
            event_data = json.loads(event.data)
            events.append(event_data)
            print(json.dumps(event_data, indent=2))
            elapsed_seconds = (datetime.datetime.utcnow() - start_time).total_seconds()
            if elapsed_seconds >= duration_seconds:
                break

    with open(output_file, 'w') as f:
        json.dump(events, f, indent=2)

    return events

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Stream events and save the output")
    parser.add_argument("--subscription_uuid", required=True, help="Subscription UUID")
    parser.add_argument("--output_file", required=True, help="Output file to save the events")
    parser.add_argument("--duration_seconds", required=True, type=int, help="Duration in seconds to stream events")

    args = parser.parse_args()

    pf_token = os.environ.get("PF_TOKEN")
    pf_host = os.environ.get("PF_HOST")

    if not pf_token or not pf_host:
        print("PF_TOKEN and PF_HOST environment variables are required.")
        exit(1)

    try:
        print(f"Starting stream at {datetime.datetime.utcnow().isoformat()}")
        stream_thread = threading.Thread(target=consume_client, args=(pf_token, pf_host, args.subscription_uuid, args.output_file, args.duration_seconds))
        stream_thread.start()
        stream_thread.join(args.duration_seconds)
    finally:
        print(f"Ending stream at {datetime.datetime.utcnow().isoformat()}")
        logout(pf_token, pf_host)
