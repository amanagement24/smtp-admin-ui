**Refactor**

- all methods Render* have a parameter that is of a type called View*
- View* variables are filled 9% with stuff taken from the SessionStore
- Of course there are other things that are captured somewhere else, not from the store. 

1. Identify all those types
2. For each type, create a method of SessionStore that, with data from the SessionStore itself AND with a number of parameters passed separately, for data that is not available in the store, returns an object of the View* type
3. See where those objects are instantiated and replace all occurrences with calls to the newly created methods

