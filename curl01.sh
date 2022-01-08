#!/bin/bash
curl -i -XPOST 'http://localhost:8086/write?db=our_expenses' --data-binary @file
