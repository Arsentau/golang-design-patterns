# [Bridge](bridge.go)

**Bridge** is a structural design pattern that lets you split a large class or a set of closely related classes into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.

## Pros

- You can create platform-independent classes and apps.
- The client code works with high-level abstractions. It isn’t exposed to the platform details.
- Open/Closed Principle. You can introduce new abstractions and implementations independently from each other.
- Single Responsibility Principle.
- You can focus on high-level logic in the abstraction and on platform details in the implementation.

## Cons

- You might make the code more complicated by applying the pattern to a highly cohesive class.
- The pattern itself incurs in a violation of the open close principle.
