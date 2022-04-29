#!/usr/bin/env python
# -*- coding: utf-8 -*-

from matplotlib import pyplot as plt
import numpy as np
from matplotlib.backends.backend_pdf import PdfPages
import sqlite3
import sys, os

# setup commandline flags
if len(sys.argv) > 2 and not 0 < int(sys.argv[2]) < 100:
    print("ERROR: Perentile has to be between 0 and 100", file = sys.stderr)
    sys.exit(1)

try:
    dbpath = sys.argv[1]
except:
    dbpath = './cs3probes/data/logs/logs.db'
try:
    perc = int(sys.argv[2])
except:
    perc = 90


# connect to database, get available probe-tables
sqliteConnection = sqlite3.connect(dbpath)
sql_query = """SELECT name FROM sqlite_master
  WHERE type='table';"""
cursor = sqliteConnection.cursor()
cursor.execute(sql_query)
tables = [v[0] for v in cursor.fetchall() if v[0] != ("sqlite_sequence") if v[0] != ("targets")]


# write data contained in tables to hashmap data
data = {}

for t in tables:
    sqliteConnection.row_factory = sqlite3.Row
    sql_query = "select * from "+t+";"
    cursor.execute(sql_query)
    a = list(cursor.fetchall())
    names = [description[0] for description in cursor.description]
    cols  = {}
    for o in names:
        cols[o] = []
    for o in a:
        for l in range(len(o)):
            cols[names[l]].append(o[l])

    data[t] = cols


# create ouput folder if it does not exist
if not os.path.exists('out'):
    os.makedirs('out')

# for each probe, calculate p-percentile and generate histogram
for probe in data.values():
    for column in probe:
        if column != "timestamp" and column != "id" and column != "targetId" and column != "portscan":
            p90 = np.percentile(probe[column], perc)
            print("\nGenerating Histogram for Test " + column)
            marks = np.array(probe[column])
            fig, axis = plt.subplots(figsize =(10, 5))
            b = np.arange(0, max(probe[column])+min(probe[column]), 1).tolist()
            axis.hist(marks, bins = b)
            with PdfPages('./out/p' + str(perc) + '_' + column + '.pdf') as export_pdf:
                plt.title('Antwortsverteilung für Test ' + column + ', n = ' + str(len(probe[column])))
                plt.xlabel('Antwortzeit in ms')
                plt.ylabel('Anzahl')
                plt.axvline(label='0-Perzentil', x=min(probe[column]), color='green', linestyle='-')
                plt.axvline(label='100-Perzentil', x=max(probe[column])+1, color='green', linestyle='-')
                plt.axvline(label=str(perc)+'-Perzentil', x=p90+1, color='red', linestyle=':')
                plt.legend(loc="upper right")
                export_pdf.savefig()
                plt.close
                print("...saved.")
