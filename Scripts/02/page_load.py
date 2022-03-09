#!/usr/bin/python3
####################################################
# Calculate the website's loading time with Python.#
####################################################
import urllib.request as urllib2
from time import time
import sys
hostlen = (len(sys.argv))

if hostlen == 2:
    host = sys.argv[1]
    stream = urllib2.urlopen(host)
    start_time = time()
    output = stream.read()
    end_time = time()
    stream.close()
    print("Time taken to load the page", host, round(end_time-start_time, 3))
    #print ("Time taken to load the page", host)
    #print(round(end_time-start_time, 3))
else:
    print ("usage: ./page_load.py <domainname>")
    print ("example: ./page_load.py https://google.com")
