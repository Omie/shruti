#!/usr/bin/python

import time
import requests
import feedparser

visitedStories = {}

def doWork():

    while True:
        feed = feedparser.parse('http://feeds.bbci.co.uk/news/world/rss.xml?edition=uk')
        for entry in feed['entries']:
            visited = visitedStories.get(entry["id"], False)
            if not visited:
                _data = { "text" : entry['title'],
                          "voice": "Salli",
                          "provider": "bbc"}
                r = requests.post("http://127.0.0.1:9574/push/", data=_data)
            visitedStories[entry["id"]] = True
        time.sleep(15 * 60) # ~15minutes
        print "bbc: time to check updates again"

if __name__ == "__main__":
    doWork()


