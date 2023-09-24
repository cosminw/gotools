package queue

import "log"

type Queue []interface{}

func(q *Queue) Push(element interface{}){
	// append's element to queue
	*q = append(*q, element)
}

func(q *Queue) PopIdx(i int)interface{}{
	//returns poped element and removes it from queue
	qu := *q
	el := qu[i]
	*q = append(qu[:i], qu[i+1:]...)
	log.Printf("poped: %+v", el)
	return el
}

func(q *Queue)Get()interface{}{
	//  retruns 1 element and removes it from queue
	qu := *q
	if len(qu) >0{
		el := qu[0]
		*q = qu[1:]
		return el
	}
	if len(*q) == 0 {
	log.Printf("empty queue")
	}
		
	return nil
}
func(q *Queue) Size()int{
	// retuns the size of queue
	return len(*q)
}

func NewQueue() *Queue{
	// NewQueue implements a queue 
	return &Queue{}
}
