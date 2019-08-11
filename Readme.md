# Flamingo Example Openweather

In this example, we build a module with connection to an api - step by step.

**Goal**: Opening `weather/wiesbaden` should show a nice weather widget.

For this example, you should be already familiar with the [Flamingo basics](https://docs.flamingo.me/) and the [helloworld tutorial](https://github.com/i-love-flamingo/example-helloworld).

## Module planning

We want to use the ports and adapters pattern for our module

| Layer | Implementations
|-------|-----------------
| Infrastructure | A fake implementation and the real implementation against openweather
| Interfaces | Web controller to show weather
| Application | Simple Service as entry
| Domain | Simple model to represent weather information

## Getting started

```bash
git clone git@github.com:i-love-flamingo/example-openweather.git
cd example-openweather
git checkout start-over
make frontend
make serve
open http://localhost:3322/
```

The start-over branch provides a skeleton Flamingo project to start with. Following the tutorial steps, you will reach
the status of the master branch. 

> Note: While you are working you can activate the carotene file watcher via `make frontend-dev`.

### Step 1 - Kickstart weather module

* Create new Flamingo module in `src/openweather`
* Create a controller that renders the page `weather/wiesbaden`, with `wiesbaden` interpreted as parameter
* Register a route using a Flamingo routes module and a (GET) handler

### Step 2 - Weather domain

Define a useful domain model.

 It's ok to conform with the api (subset) [openweathermap.org/current]( https://openweathermap.org/current)

  We will use the domain model to be passed to the template later.

Start with mocking the template variable and display the weather by using the `debug` mixin:

```pug
include /atom/debug/debug

// ...

+debug(weather)
```

### Step 3 - Fake service

* Implement backend logic inside the module -  following ports and adapters.
* Plan it: What needs to be done in which layer?  (start from inner to outer)?
  * Define a `domainService` (as secondary port)
  * Add `applicationService` (that gets the adapter injected)
  * Add a `fakeImplementation` for the `domainService`
  * Bind the fake implementation to the `domainService` in your module's `Configure` function
  * Call the `applicationService` from the controller
  
### Step 4 - Real service

* Implement an adapter against the openweather API and adjust the template to show nice HTML instead of debug output.
  
  Use `http://openweathermap.org/img/w/${weather.iconCode}.png` for a nice image.
  
  Push the variable content to the browser console using the `console` mixin:
  ```pug
  include /atom/console/console
  
  // ...
  
  +console(weather)
  ```
* We suggest to use pact for consumer functional testing of api implementations
  * For further reading: [docs.pact.io](https://docs.pact.io/) 
  * Usage of Flamingo’s module `core/testutil` to support pact tests
  * Setup Pact on your dev host: [github.com/pact-foundation/pact-go](https://github.com/pact-foundation/pact-go)

**You can start into this step with the skeleton provided in `step_4_start.patch`:**
```bash
git apply step_4_start.patch
``` 


#### Typical adapter implementations

1. Call the external API 
1. Unmarshal Response  to DTO (DataTransferObjects) 
1. Mapper DTO -> DomainModel (Anti Corruption Layer) 

### Step 5 - Data controller

Add a data controller to show the current weather of „Wiesbaden“ in the Header by using Flamingo pugtemplate’s template function `data`.

### More time? Improve the example!

Ideas:
* Use „locale“ package to translate „mainCharacteristics“
* Add template function to convert kelvin to celcius
* Add form to select the city
* …
