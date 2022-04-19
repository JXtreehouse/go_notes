
// java中的future机制
// 当启动一个函数或者执行一个task时候，并不需要马上知道他的结果，当需要知道这个结果时候才去get这个结果(异步返回)
// 在这个期间可以执行一些其他的逻辑
// 当get这个结果时候如果结果还没出来，就会阻塞在那里，如果结果已经出来，就能get到，然后得到这个结果以后继续往下执行，这种机制可以大大提升程序的效率。
// https://www.youtube.com/watch?v=5g7C4A2ErkI
private static FutureTask<String> service() {
    FutureTask<String> task = new FutureTask<String>(()-> "Do someting")
    new Thread(task).start()
    return task;
}

FutureTask<string> ret = service();
    System.out.println("Do someting else");
    System.out.println(ret.get());