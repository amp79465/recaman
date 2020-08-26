import sys, time

#Unpack the sequence object and see if the number is in the range.
def inSequence(n):
    for set in rMembers:
        if n >= set[0] and n <= set[-1]:
            return True
    return False

#Add the number to the sequence object.
def addN(n):
    for i in range(0, len(rMembers)):
        #Handle if we are at the end of the list
        if i == len(rMembers) - 1:
            if n > rMembers[i][-1] + 1:
                rMembers.append([n])
                return
            if n == rMembers[i][-1] + 1:
                rMembers[i] = [rMembers[i][0], n]
                return
        #If we aren't close enough to change anything, pass
        elif n < rMembers[i-1][0] or n > rMembers[i+1][-1]:
            pass
        #Handle adding the number on the left side of the current index
        elif n < rMembers[i][0]:
            if n > rMembers[i-1][-1] + 1 and n == rMembers[i][0] - 1:
                rMembers[i] = [n, rMembers[i][-1]]
                return
            elif n == rMembers[i][0] - 1 and n == rMembers[i - 1][-1] + 1:
                rMembers[i] = [rMembers[i-1][0], rMembers[i][-1]]
                del rMembers[i-1]
                return
            elif n == rMembers[i - 1][-1] + 1 and n < rMembers[i][0] - 1:
                rMembers[i-1] = [rMembers[i-1][0], n]
                return
            elif n < rMembers[i][0] - 1 and n > rMembers[i-1][-1] + 1:
                rMembers.insert(i, [n])
                return
        #Handle adding the number on the right side of the current index
        elif n > rMembers[i][-1]:
            if n < rMembers[i+1][0] - 1 and n == rMembers[i][-1] + 1:
                rMembers[i] = [rMembers[i][0], n]
                return
            elif n == rMembers[i][-1] + 1 and n == rMembers[i+1][0] - 1:
                rMembers[i] = [rMembers[i][0], rMembers[i+1][-1]]
                del rMembers[i+1]
                return
            elif n == rMembers[i+1][0] - 1 and n > rMembers[i][-1] + 1:
                rMembers[i+1] = [n, rMembers[i+1][-1]]
                return
            elif n > rMembers[i][-1] + 1 and n < rMembers[i+1][0] - 1:
                rMembers.insert(i+1, [n])
                return

#Variable declarations, first argument on command line is termLimit
termLimit = int(sys.argv[1])
rMembers = [[0]]
lastAdded = 0

#Main loop of the program, the output file being the second argument of the file
with open(sys.argv[2], 'w') as outFile:
    outFile.write("0,0\n")
    beginTime = time.perf_counter()
    #The main sequence calculation
    for n in range(1, termLimit + 1):
        rN = lastAdded - n
        if rN > 0 and not inSequence(rN):
            addN(rN)
            outputString = str(n) + "," + str(rN) + "\n"
            outFile.write(outputString)
            lastAdded = rN
        else:
            addN(lastAdded + n)
            outputString = str(n) + "," + str(lastAdded + n) + "\n"
            outFile.write(outputString)
            lastAdded = lastAdded + n
    endTime = time.perf_counter()
#Write out the object to file for inspection
fileName = "rMembersObject-" + str(sys.argv[2])
with open(fileName, 'w') as fileOut:
    for member in rMembers:
        fileOut.write(str(member))
print("Length of sequence object: ",len(rMembers))
print(str(len([m for m in rMembers if len(m) == 1])), " single digit members")
print(str(len([m for m in rMembers if len(m) == 2])), " double digit members")
print(f"Completed in {endTime - beginTime:0.4f} seconds")
