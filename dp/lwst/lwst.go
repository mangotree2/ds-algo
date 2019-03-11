package lwst

//莱文斯坦距离（Levenshtein distance）
//如果：a[i]!=b[j]，那么：min_edist(i, j)就等于：
//min(min_edist(i-1,j)+1, min_edist(i,j-1)+1, min_edist(i-1,j-1)+1)
//
//如果：a[i]==b[j]，那么：min_edist(i, j)就等于：
//min(min_edist(i-1,j)+1, min_edist(i,j-1)+1，min_edist(i-1,j-1))
//
//其中，min表示求三数中的最小值。

//回溯是一个递归处理的过程。如果a[i]与b[j]匹配，我们递归考察a[i+1]和b[j+1]。如果a[i]与b[j]不匹配，那我们有多种处理方式可选：
//
//可以删除a[i]，然后递归考察a[i+1]和b[j]；
//
//可以删除b[j]，然后递归考察a[i]和b[j+1]；
//
//可以在a[i]前面添加一个跟b[j]相同的字符，然后递归考察a[i]和b[j+1];
//
//可以在b[j]前面添加一个跟a[i]相同的字符，然后递归考察a[i+1]和b[j]；
//
//可以将a[i]替换成b[j]，或者将b[j]替换成a[i]，然后递归考察a[i+1]和b[j+1]。


