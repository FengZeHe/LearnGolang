### 队列
队列是一种特殊的线性表，它智能在表的前端（front）进行删除操作，
而在表的尾端(rear)进行插入操作。进行插入操作被称为尾端，进行删除
操作端成为队头。

#### 顺序队列 ＆ 链式队列
队列可以用数组和链表实现。用数组实现的队列叫做顺序队列，用链表实现的
队列叫做链式队列。队列需要两个指针，一个是head指针用来指向队头，
一个是tail指针用来指向队尾。

#### 顺序队列 ＆ 链式队列优缺点
1. 使用数组做队列优点：可以随机访问元素并更改
2. 使用数组作队列缺点：在数组中插入元素后，对之后的元素都要进行数据搬移
3. 使用链表做队列优点：插入或修改元素方便，只需改变指针指向，无需数据搬移
4. 使用链表做队列缺点：随机访问元素需要遍历链表



