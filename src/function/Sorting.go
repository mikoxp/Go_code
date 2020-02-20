package function



func BubbleSort(tab []float32) []float32{
	n:=len(tab)
	for i:=0; i<n-1;i++{
		for j:=0;j<n-1;j++{
			if tab[j]>tab[j+1]{
				tmp:=tab[j]
				tab[j]=tab[j+1]
				tab[j+1]=tmp
			}
		}
	}
	return tab
}