# Goxide Samples.
Use this API to generate JSON Sample data for your test projects. Information for Users, Products,Todos, Posts and Comments. 

## Users
You can generate up to 100 User data per request. by default, you'll have 10 users. you may pass a total query parameter to your call.
User ID is used across all data and can be used to get information for all users in all other fields except products.
 
### Usage Examples

```
goxide.com/users?total=35
```

``` Sample User Data ``` 
{
    
}

## Todos
You get a default of 10 todos on a single call.
``` Sample Todo Data ``` 
{

}

## Posts 
Get a  default of 20 posts on a single call. you can use the ?total query to specify the number of posts needed.
``` Sample Post Data ``` 
{

}

## Comments
Get General comments or pass comments with a specific post id to get comments under a single posts.
``` Sample Comment Data ``` 
{

}

## Products

Refer to the side navigation for http methods available for all information.