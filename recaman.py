import sys, time

termLimit = int(sys.argv[1])
rMembers = [0]
lastAdded = 0

beginTime = time.perf_counter()
for n in range(1, termLimit + 1):
    rN = lastAdded - n
    if rN > 0 and rN not in rMembers:
        rMembers.append(rN)
        lastAdded = rN
    else:
        rMembers.append(lastAdded + n)
        lastAdded = lastAdded + n
endTime = time.perf_counter()
print(f"Completed in {endTime - beginTime:0.4f} seconds")
