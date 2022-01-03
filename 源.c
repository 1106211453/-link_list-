#include <stdio.h>
#include <stdlib.h>
#include "a.h"
/*typedef struct _note {
	int value;
	struct _note* next;
}Note;*/
typedef struct _list {
	Note* head;
}List;
void add(List* plist, int number);//创造链表
void print(List* plist);//打印链表
void listInsert(List* list, int i, int e);//给链表插入一个节点和一个数e,节点插到i右边
int main()
{
	List list;
	list.head = NULL;
	int number=0;
	int d = 0;
	do
	{
		scanf_s("%d", &number);
		if (number != -1)
		{
			add(&list, number);
		}
	} while (number!=-1);
	print(&list);
	printf("请输入要找的数\n");
		scanf_s("%d", &number);
	int isfonud = 0;
	Note* p;
	for (p=list.head;p;p=p->next)
	{
		if (p->value == number)
		{
			printf("找到了%d\n", p->value);
			isfonud = 1;
		}
	}
	if (isfonud == 0)
	{
		printf("没找到");
	}
	Note* q;
	for (q=NULL,p=list.head;p;q=p,p=p->next)
	{
		if (p->value == number)
		{
			if (q) {

				q ->next= p->next;
			}
			else
			{
				list.head = p->next;
			}
			free(p);
			break;
		}
	}
	print(&list);
	printf("请输入要插入的数\n");
	scanf_s("%d", &d);
	listInsert(&list, 9, d);
	print(&list);
	for (p=list.head;p; p = q)
	{			
		q = p->next;
		free(p);
	}
	return 0;
}
void add(List *plist,int number)
{
	Note* p = (Note*)malloc(sizeof(Note));
	p->value = number;
	p->next = NULL;
	Note* last = plist->head;
	if (last)
	{
		while (last->next) 
		{
			last = last->next;
		}
		last->next = p;
	}
	else
	{
		plist->head = p;
	}
}
void print(List* plist)
{
	Note*p;
	for (p=plist->head; p;p= p->next)
	{
		printf("%d\t",p->value);
	}
	printf("\n");
}
void listInsert(List* list, int i, int e)//少了一个，i不存在的情况(已修改)
{
	int j = 0;
	Note* s = (Note*)malloc(sizeof(Note));
	Note* p;
	for (p = list->head;j<i-1;p=p->next)
	{
		if (p ->next== NULL)
		{
			break;
		}
		j++;
	}
	s->value = e;
	s ->next= p->next;
	p->next = s;
}