# Graph Algorithms Playbook

Graph questions explore traversal, shortest paths, and connectivity. This folder groups templates and interview-style challenges so you can adapt quickly to new constraints.

## Interview Focus Areas

- Traversal mastery: BFS for level-order insight, DFS for component discovery, and iterative vs recursive tradeoffs.
- State tracking: visited sets, parent maps, and coloring strategies for detecting cycles or bipartite graphs.
- Graph representations: adjacency lists vs matrices, edge lists, and when to reach for each in Go.
- Topological reasoning: ordering jobs, evaluating dependencies, and spotting DAG structures in disguise.

## Real-World Applications

- Navigation systems: routing through cities, road networks, and transit timetables.
- Dependency management: package builds, task schedulers, and CI/CD pipelines.
- Social and recommendation graphs: community detection, friend suggestions, and influence propagation.

## What's in This Folder

- `bfs.go`: level-order traversal template that generalizes to multi-source BFS problems.
- `dfs.go`: recursive DFS helper for connected components and cycle detection.
- `representation.go`: adjacency list structures to jump-start custom graph shapes.
- `questions/`: interview-formatted problems like city edge reordering, island counting, and path existence checks.
- `resources.md`: curated learning links covering theory refreshers and advanced scenarios.

## Practice Game Plan

1. Re-implement BFS and DFS iteratively and recursively to ensure fluency.
2. Tackle `questions/number_of_islands.go` and variants to reinforce grid-to-graph translation.
3. Simulate topological sorting by altering edge directions and verifying cycle detection logic.
4. Add weighted edges and experiment with Dijkstra or Bellman-Ford to round out shortest-path coverage.

## Stretch Topics

- Study Union-Find (Disjoint Set Union) for dynamic connectivity and Kruskal's MST.
- Explore bidirectional search and A* for heuristic-driven traversals.
- Model real interview prompts from LeetCode hard/Graph category and compare approaches.
