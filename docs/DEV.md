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

To get the content (CORS issue prevents loading it directly via htmx)

```
mkdir -p content && cd content
nix-shell -p httrack
```
`httrack "https://nightly.changelog.com" -O . "-*email*" --preserve`

## TODO

- [x] htmx ui
    - [ ] select only the specific elements from the server
    - not possible to select multiple things from response
    - I guess oob swap or client side scripting
- [x] client side caching to avoid requests
    - see if fly.io functions like this already without doing much
    - assuming this works because fly.io is sophisticated enough for static assets
- [ ] separate categories for first timers, top new, repeat performers?
    - would make more sense to do this from db rendered ui not static html files parsed
- [ ] calendar view with htmx/cache as the simplest approach
    - [ ] Better calendar on desktop
        - [ ] show skull/question mark icons in known missing entries in the calendar
    - [ ] mobile ui is enough (bigger bottom panel)
- [ ] missing issues detect via 404 and also 404 page as well
- [x] prevent <2015-01-01 and >=current date in calendar ui
- [ ] light dark mode switch
    - how is it done in the email? where is this info saved?
    - I guess client side is enough
    - no cookies, avoid gdpr cookie prompt
- [x] google analytics
    - not sure, will require input from the changelog people
- [x] Should I be writing this in ruby?
    - Intent is to learn htmx with go (learn templ later)
- mistyped as nightlog.changley.com
    - once in while this can be shown from the server side (P<0.02) and would act as an easter egg
    - not going to implement it
- [x] left right buttons (not naviagtion but date wise)
    - [x] keyboard shortcuts
    - [x] Bug with sholeace icons requests
    - [x] disable hx-get for empty urls
- [ ] Vendor all deps
    - Client side
    - [x] htmx
    - [x] htmx extenstions
    - [ ] shoelace (is this possible)
        - [ ] bun install and serve?
    - [ ] embed into go binary, static scripts
- [ ] optional memcache/varnish/redis caching
    - try this with lru eviction or something to speed up
    - what about today()? it will require cache invalidation when changed
- [ ] remove all todos and the todo section once done
- [x] open issue for discussion https://github.com/thechangelog/nightly/issues/44
- [ ] send pr
    - [ ] remove proxying glue code/endpoints proxying nightly.changelog.com
