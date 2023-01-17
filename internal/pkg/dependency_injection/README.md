# Dependency Injection

There are two main ways to get it done in practice. The first one is manually, and the other (and prettier) is using a dependency injection container.

## [Manually](manual/manual.go)

Manually construction is an objective way to do it. You declare, create, and inject your dependencies step by step. I think it’s clean and there isn’t any magic happening behind the scenes. The problem is as your dependencies get complex you need to deal with complexity by yourself. You may see your func main() getting with hundreds of lines of code and harder to maintain.

## Containers (TODO)

In a container style, you’ll need to teach your container how to create a dependency and then it’ll create your dependency graph to discover how to create dependencies. Once you ask for a dependency, it’ll follow the graph creating everything related to it. Let’s imagine a scenario where we need to create two Services/Use cases, and they have dependencies between them and across them.

Once you've taught your container how to create each dependency, it'll create a dependency graph.

f you’ve chosen the manual style, imagine this scenario getting bigger and bigger. Imagine all of your services doing the same startup process. It’ll get complex to deal with in long term.

### [uber-go/dig](https://github.com/uber-go/dig)

> is a dependency injection toolkit developed by Uber to resolve this kind of problem, with reflection. It’s really powerful and helps you to reduce the code you write to setup your service. It’s important to say it’s focused on the application startup. By design, uber-go/dig understands that everything is a Singleton and it'll create your dependencies just once.

### [sarulabs/di](https://github.com/sarulabs/di)

> Dependency injection framework for go programs (golang).
> DI handles the life cycle of the objects in your application. It creates them when they are needed, resolves their dependencies and closes them properly when they are no longer used.
>
>If you do not know if DI could help improving your application, learn more about dependency injection and dependency injection containers:
>
>What is a dependency injection container and why use one ?
>There is also an Examples section at the end of the documentation.
>
>DI is focused on performance. It does not rely on reflection.
