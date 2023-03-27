local list = nil
local head = list

local list = {next = list, val=1}
local list = list.next
local list = {next = list, val=2}
local list = list.next
local list = {next = list, val=3}
local list = list.next
local list = {next = list, val=4}
local list = list.next
local list = {next = list, val=5}

local l = head
while l do 
	print(l.val)
	l = l.next
end
