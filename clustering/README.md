# Clustering

## Algorithm Pseudo-code

- Load 1000 points
- For each point
  - If the point fits in a cluster in the discard set, add it
  - If the point fits in a cluster in the compressed set, add it
  - Add the point to the retained set
- For each pair of points in the retained set, if they can merge, merge them and add to the compressed set
  - If the std deviation of the resulting set would be less than 4, merge
- Go through discard sets and merge them if their merged has a mahalonobis distance <= 4d