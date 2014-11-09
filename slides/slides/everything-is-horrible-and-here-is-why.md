##  Everything is horrible and here is why

Here are some common issues with the languages listed:

- Lack of concurrency primitives which leads to Callback hell/bringing in thrid party libs and frameworks
- For Python/Ruby the gil makes true threading/parrell tasks less then useful
- Deployment is a lot harder than it needs to be (Virtualenvs, Gems, Node_modules), Is the package server (pip, npm, gem) up and running, Is the correct package there, is build on the same os?
- Lack of type safety causes high bug rate
- While node can be very fast at IO, python, Ruby, and Node all fall down when it comes to cpu and/or ram usage
