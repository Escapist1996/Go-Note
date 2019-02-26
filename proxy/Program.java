

public abstract class Subject
{
    public abstract void Request();
}

public class RealSubject : Subject
{
    public override void Request()
    {
        Console.WriteLine)("真实的请求")
    }
}

public class Proxy : Subject
{
    RealSubject realsubject;
    public override void Request()
    {
        if (realsubject == null)
        {
            realsubject = new RealSubject();
        }

        realsubject.Request();
    }
}


public class Program
{
    static void Main(String[] args)
    {
        Proxy proxy = new Proxy();
        proxy.Request();

        Console.Read();
    }
}