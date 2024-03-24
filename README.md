https://nightly.changelog.com/

- run missing.js contents in the inspector of above url
- missing.js script made using chatgpt

### Output as of 2024/03/24

```
Missing dates for 2015:
Missing dates for 2016:
Missing dates for 2017:
Missing dates for 2018:
2018-07-10
2018-07-11
2018-07-12
Missing dates for 2019:
Missing dates for 2020:
2020-08-22
Missing dates for 2021:
2021-05-08
2021-05-10
2021-05-11
2021-08-25
2021-08-26
2021-08-27
2021-10-06
2021-10-22
2021-10-24
2021-10-25
2021-10-26
2021-10-27
2021-10-28
Missing dates for 2022:
2022-12-06
Missing dates for 2023:
Missing dates for 2024:
```

## TODO

- [ ] htmx ui and select only the specific elements from the server
- [ ] client side caching to avoid requests
    - see if fly.io functions like this already without doing much
- [ ] separate categories for first timers, top new, repeat performers?
- [ ] calendar view with htmx/cache as the simplest approach
- [ ] missing issues detect via 404 and also 404 page as well
- [x] prevent <2015-01-01 and >=current date in calendar ui
- [ ] light dark mode switch
    - how is it done in the email? where is this info saved?
    - I guess client side is enough for us
    - no cookies
- [ ] google analytics
    - not sure, will require input from the changelog people
- [x] Should I be writing this in ruby?
    - Intent is to learn htmx with go (learn templ later)
- [ ] remove all todos and the todo section once done

