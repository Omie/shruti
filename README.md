shruti
------

listen to all those updates using Ivona TTS

This is a work-in-progress. In PRE PRE PRE ALPHA stage.
A lot of things are expected to be changed.


How it works
-------------
- shruti is a simple server that keeps listening for new TTS requests
- various other programs do the labour. Fetch updates, scrape, consume api or
 whatever. Then they push their notifications to shruti
- shruti will get the TTS part done using Ivona and push it to its queue.
- After specified interval, shruti will make sure to read out those notifications
- one after another


Why
----
- I am bored of reading everything myself. I would love if someone does the job for me
- Something that will check
    - twitter. My timeline and few accounts I like to keep up with
    - HN, few reddit channels
    - some headlines from bbc
    - some important server statuses
    - build server status updates if important
    - New email upon arrival. From and Subject
    - Share market updates. Few important share values I care about
    - any any other sites I can think of
- ... and read out loud every 15-30 minutes.
- Hence (:


About name
-----------
- The shruti is a Sanskrit term, which in the context of Indian music, is considered the smallest interval of pitch that the human ear can detect
- [Shruti - wiki](http://en.wikipedia.org/wiki/Shruti_%28music%29)
- besides, its a girl's name, suits the voice and sounds catchy enough


How to write a provider
-----------------------
- Use any language, do whatever to get the text
- Send an HTTP request to shruti server at `http://host:port/push/`
    - parameters required:
        - `text`  : string : text to be spoken
        - `voice` : string    : ivona voice [not implemented yet]
        - `provider` : string : provider name for reference
        - `keepfile` : boolean : yes/no to keep file [not implemented yet]

How to set up
-------------
- sign up for beta account on ivona.com
- generate credentials
- install ffmpeg. We need ffplay binary
- get the shruti server binary, make sure to keep it up and running
- make sure to export environment variables `SHRUTI_ACCESSKEY` and `SHRUTI_SECRETKEY`
- make sure you have `audiofiles` directory created
- run your providers however you want
- You could use supervisor or upstart on Ubuntu.

Goal
-----
- Build an open source SaaS that anyone can deploy for themselves, use it personally
- Do not go after building J.A.R.V.I.S. There is already a project for that[jasper](http://jasperproject.github.io/)

- TODO:
    - Convert this hack into a nice future proof API, make it configurable easily
    - Support multiple audio player backends to make it usable on Windows and Mac
    - Add providers
    - Add actions, Mute, Unmute, Force speak
    - Build configuration, status pages
    - Make use of different voices for different providers
    - Add support for noticing special tags probably. Add extra text to such texts to make it more natural

Contributions, thoughts are welcome :)

License MIT


