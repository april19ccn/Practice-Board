文档：https://docs.github.com/zh/rest/search/search?apiVersion=2022-11-28#search-issues-and-pull-requests
q字段：https://docs.github.com/zh/search-github/searching-on-github/searching-issues-and-pull-requests#search-within-a-users-or-organizations-repositories

Graphql 接口（也是github页面使用的）：使用的_graphql，可能包含一些隐藏条件，故数据与API的数目不一致（即使api加上is:issue API会多拿一些）
https://github.com/_graphql?body={%22query%22%3A%2229746fd23262d23f528e1f5b9b427437%22%2C%22variables%22%3A{%22name%22%3A%22go%22%2C%22owner%22%3A%22golang%22%2C%22query%22%3A%22state%3Aopen%20json%20decoder%20sort%3Acreated-desc%22}}

REST 接口：
https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+state%3Aopen+json+decoder+is%3Aissue&sort=created&order=desc&per_page=&page=

按创建时间获取：（可以在页面直接书写测试之后粘贴过来）
created%3A>2025-05-01
（repo%3Agolang%2Fgo+is%3Aopen+json+decoder+created%3A>2025-05-01）


## Issue 操作
- 创建一个 issue
https://docs.github.com/zh/rest/issues/issues?apiVersion=2022-11-28#get-an-issue