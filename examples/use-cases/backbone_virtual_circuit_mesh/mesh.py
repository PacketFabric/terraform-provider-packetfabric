from itertools import combinations

#sample_list = ['a', 'b', 'c']
#sample_list = ['PF-AP-SFO13-1748214', 'PF-AP-MIA3-1715317', 'PF-AP-LAS1-1715316', 'PF-AP-LAX1-8620', 'PF-AP-NYC1-824', 'PF-AP-DAL1-823']
sample_list = ['var.pf_port1', 'var.pf_port2', 'var.pf_port3', 'var.pf_port4', 'var.pf_port5', 'var.pf_port6']
list_combinations = list()
sample_set = set(sample_list)

for n in range(len(sample_set) + 1):
    list_combinations += list(combinations(sample_set, n))

#print(list_combinations)

s=0
length = len(list_combinations)
for i in range(length):
    if len(list_combinations[i]) == 2:
       print(list_combinations[i]) 
       s = s +1

print("\nMesh size: " + str(s))