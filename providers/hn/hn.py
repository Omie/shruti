#!/usr/bin/python

import time
import requests


visitedStories = {}

def doWork():

    while True:
        _resp = requests.get('https://hacker-news.firebaseio.com/v0/topstories.json').json()
        for storyid in _resp[:30]:
            visited = visitedStories.get(storyid, False)
            if not visited:
                _resp = requests.get('https://hacker-news.firebaseio.com/v0/item/{0}.json'.format(storyid)).json()
                _data = { "text" : _resp['title'],
                          "voice": "Salli",
                          "provider": "hackernews"}
                r = requests.post("http://127.0.0.1:9574/push/", data=_data)
            visitedStories[storyid] = True
        time.sleep(15 * 60) # ~15minutes
        print "time to check updates again"

if __name__ == "__main__":
    doWork()


