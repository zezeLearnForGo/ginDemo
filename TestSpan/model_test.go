package TestSpan

import (
	"fmt"
	"ginDemo/model"
	"testing"
)

func TestName(t *testing.T) {

	var spanList []main.Span

	guiref := main.Ref{
		TraceId:      "1",
		SpanId:       "22",
		ServiceName:  "webapp",
		Relationship: "child",
	}

	guiref2 := main.Ref{
		TraceId:      "1",
		SpanId:       "22",
		ServiceName:  "webapp",
		Relationship: "child",
	}

	var refList []main.Ref

	refList = append(refList, guiref)
	refList = append(refList, guiref2)

	webapp := main.Span{
		TraceId:     "1",
		SpanId:      "22",
		ServiceName: "webapp",
	}

	gui := main.Span{
		SpanId:      "2",
		TraceId:     "2",
		ServiceName: "gui",
		Ref:         refList,
	}

	// webapp - text
	var refListText []main.Ref

	textref := main.Ref{
		TraceId:      "1",
		SpanId:       "22",
		ServiceName:  "webapp",
		Relationship: "child",
	}
	refListText = append(refListText, textref)

	text := main.Span{
		SpanId:      "3",
		TraceId:     "3",
		ServiceName: "text",
		Ref:         refListText,
	}

	//  gui -> webapp
	webapp1ref := main.Ref{
		TraceId:      "1",
		SpanId:       "22",
		ServiceName:  "gui",
		Relationship: "child",
	}
	var refListWebapp1 []main.Ref
	refListWebapp1 = append(refListWebapp1, webapp1ref)
	webapp1 := main.Span{
		TraceId:     "1",
		SpanId:      "2233",
		ServiceName: "webapp",
		Ref:         refListWebapp1,
	}

	spanList = append(spanList, webapp)
	spanList = append(spanList, gui)
	spanList = append(spanList, text)
	spanList = append(spanList, webapp1)

	// 除了rootSpan  每个span都有父亲  谁调用它就是它的父亲
	var RelationshipList []main.Relationship

	for i := range spanList {
		// root span
		if len(spanList[i].Ref) == 0 {
			// root span 没有父亲
			ParentName := ""
			ChildName := spanList[i].ServiceName
			if len(RelationshipList) == 0 {
				// 没有就添加
				RelationshipList = append(RelationshipList, main.Relationship{
					Parent: ParentName,
					Child:  ChildName,
					Count:  1,
				})

			}
		}

		// no root span
		if len(spanList[i].Ref) != 0 {

			// 遍历ref
			for x := range spanList[i].Ref {

				// ref
				ParentName := spanList[i].Ref[x].ServiceName
				ChildName := spanList[i].ServiceName

				if len(RelationshipList) == 0 {
					// 没有就添加
					RelationshipList = append(RelationshipList, main.Relationship{
						Parent: ParentName,
						Child:  ChildName,
						Count:  1,
					})
				}

				isExist := false
				for y := range RelationshipList {
					// 有的话count+1
					if RelationshipList[y].Parent == ParentName && RelationshipList[y].Child == ChildName {
						RelationshipList[y].Count += 1
					}
					// 没有就添加
					if RelationshipList[y].Parent != ParentName {
						isExist = false
					}
				}

				if !isExist {
					// 没有就添加
					RelationshipList = append(RelationshipList, main.Relationship{
						Parent: ParentName,
						Child:  ChildName,
						Count:  1,
					})
				}
			}
		}
	}
	fmt.Println(RelationshipList)

}
