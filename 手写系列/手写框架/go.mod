 module example

 go 1.16
 
 require turingGo v0.0.0
 
 # 在 go.mod 中使用 replace 将turingGo 指向./TuringGo
 replace turingGo => ./TuringGo