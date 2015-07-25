shruti
------


What is it
-----------

Shruti is an open source notification system that lets one listen to all those
updates instead of reading them themselves.

It uses [Ivona service](http://www.ivona.com/) to convert text to speech.

It is meant to be used as a web based service, but its trivial to use it locally as
a desktop application.

Its written in Go, builds are available for Gnu\Linux, Mac OS X and Windows.


How it works
-------------
- There are 4 major parts of the system. Backend API server, Providers, Client and
a supporting ivona microservice.
- Backend API takes care of dumping data to postgres and serve it on demand
- Providers collect data from various sources and push it to the backend server
- Client manages to read these updates at loud for you
- Ivona service makes it easier to communicate with Ivona API


btw, it also supports instant push notifications using [pusher.com](http://pusher.com)


Why
----
- I am bored of reading everything myself. I would love if someone does the job for me
- Something that will check
    - HN, few reddit channels
    - twitter. My timeline and few accounts I like to keep up with
    - some headlines from bbc
    - some important server statuses
    - build server status updates if important
    - New email upon arrival. From and Subject
    - Share market updates. Few important share values I care about
    - any any other sites I can think of
    - any other devices that keep track of something
-... and read out loud every 15-30 minutes.
- Hence (:


About the name
-----------
- "shruti" is a Sanskrit term, which in the context of Indian music, is considered the smallest interval of pitch that the human ear can detect [Shruti - wikipedia](http://en.wikipedia.org/wiki/Shruti_%28music%29)
- besides, it is quite common name over here, and sounds catchy enough!


API Specification
-----------------

Please read the spcefication at [docs.shrutiapp.apiary.io](http://docs.shrutiapp.apiary.io/)


How to write a provider
-----------------------

Please read the documentation at [shruti-providers](https://github.com/Omie/shruti-providers/)


How to set it up
-------------

Please read the documentation at [shruti-deployment](https://github.com/Omie/shruti-deployment)


Goal
-----

- Improve and be ready for the time when IoT fianlly takes over our lives with data from everywhere
- Seriously, take a look at these few
    - [ThingSpeak](https://thingspeak.com/) 
    - [xively](https://xively.com/)
    - [TEMBOO](https://www.temboo.com/)


Contributions, thoughts are welcome :)


License
-------

MIT License


