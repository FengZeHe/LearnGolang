## 图
图中的元素我们叫做顶点（vertex），图中的一个顶点可以与任意其他顶点建立连接关系。我们把这种建立关系的叫作边。

### 入度和出度
顶点的入度表示多少条边指向这个顶点；顶点的出度表示多少条边以正顶点为起点指向其他顶点。
以微博的例子表示，入度表示有多少粉丝，出度表示关注了多少人。

## 带权图
在带权图中，每条边都有一个权重(weight)，可以用这个权重表示QQ好友的亲密度。


### 深度和广度优先搜索
深度优先和广度优先搜素算法都是基于"图"这种数据结构的。因为图这种数据结构表达能力很强，大部分设计搜索到场景都可以抽象成"图"

### 广度优先搜索（BFS）
BFS是一种地毯式层层推进的搜索策略，即先查找离顶点最近的，然后次近的，依次往外搜索。

### 深度优先搜索（DFS）
