# data-structures
data structures implemented in go

Understand the pattern:

If input array is sorted then
- Binary search
- Two pointers

If asked for all permutations/subsets then
- Backtracking

If given a tree then
- DFS
- BFS

If given a graph then
- DFS
- BFS

If given a linked list then
- Two pointers

If recursion is banned then
- Stack

If must solve in-place then
- Swap corresponding values
- Store one or more different values in the same pointer

If asked for maximum/minimum subarray/subset/options then
- Dynamic programming

If asked for top/least K items then
- Heap

If asked for common strings then
- Map
- Trie

Else
- Map/Set for O(1) time & O(n) space
- Sort input for O(nlogn) time and O(1) space

Where have you seen 𝗗𝗮𝘁𝗮 𝗦𝘁𝗿𝘂𝗰𝘁𝘂𝗿𝗲𝘀 in 𝗿𝗲𝗮𝗹 𝘁𝗶𝗺𝗲 𝗔𝗽𝗽𝗹𝗶𝗰𝗮𝘁𝗶𝗼𝗻𝘀 ?

𝐒𝐭𝐚𝐜𝐤:
1) UNDO option
2) Text editor: you push letter by letter to the stack so you erase back.
3) Recursion(inbuilt stack)

𝐐𝐮𝐞𝐮𝐞:
1) Your browser deletes the history past one month.
2) If you delete a picture on your phone, it will be the "recently delete" folder which says "the images will be deleted permanently after one week".
   Here all the images are stored in the queue so it's easier to pop from the rear based on the image deletion date.
3) Waiting list: during online registrations, sometimes you'll be put on the waiting list. basically, all the requests will be stored in the queue.

𝐋𝐢𝐧𝐤𝐞𝐝 𝐥𝐢𝐬𝐭:
1) Browser's Next and Previous Button: a linked list of URLs
2) music player where you can play the next or previous song. (Doubly linked list).
3) In the ludo game, It has to pass the chance to each player in a circular fashion (circular linked list).

𝐓𝐫𝐞𝐞𝐬:
1) File system: Folders and subfolders (N-ary tree).
2) e-commerce websites : category -> subcategories -> products
3) Auto-suggestion when you google (Trie)

𝐆𝐫𝐚𝐩𝐡𝐬:
1) Uber, Ola cab booking: show nearest available cars (BFS)
2) Maven dependencies build order ( Topological sorting (DFS))
3) While booking bus/flights, you get a list of available routes.
4) In Facebook, users are considered to be the vertices and if they are friends then there is an edge running between them. Facebook’s Friend suggestion algorithm uses graph theory. Facebook is an example of an undirected graph.
