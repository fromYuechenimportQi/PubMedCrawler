# PubMedCrawler

爬取PubMed上的论文信息，通过关键词检索，将文章标题、作者、摘要、期刊、发表时间、摘要中文翻译(可选，百度翻译API)保存至word文档

Usage:
```
$ go build ./main/main.go
$ ./main -kw "R2R3 myb domain" -ss small -trans -tid YOUR_BAIDU_TRANSLATE_APP_ID -tsk YOUR_BAIDU_TRANSLATE_SECRET_KEY -out YOURFILE
```

必填参数：-kw (注：如果该参数中有空格，则需要用双引号括起来)


可选参数：-trans (无此参数则没有中文摘要翻译，点击此链接 http://api.fanyi.baidu.com/注册并认证， 使方可使用摘要翻译功能)


默认参数：-ss small -out default.docx

