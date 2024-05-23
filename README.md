![Fork GitHub Release](https://img.shields.io/github/v/release/fRead-dev/htmlValidator)
![Tests](https://github.com/fRead-dev/htmlValidator/actions/workflows/go-test.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/fRead-dev/htmlValidator)](https://goreportcard.com/report/github.com/fRead-dev/htmlValidator)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/fRead-dev/htmlValidator?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/fRead-dev/htmlValidator?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/fRead-dev/htmlValidator)


# htmlValidator

1. Абзац
    - def `<p>`
    - лево `<p left>`
    - право `<p right>`
    - центр `<p center>`
    - разделитель `<hr>`
2. Стилистика
    - жирный `<b>`
    - наклонный `<i>`
    - подчеркнутый `<u>`
    - зачеркнутый `<s>`
    - цитата `<q>`
    - в низ мелкий текст `<sub>`
    - в верх мелкий текст `<sup>`

Правила:

- в абзаце не может быть других абзацев. 
- стили могут стаковатся беконечно
- запрещены любые другие теги кроме описаных


```css
p[left]{text-align: left;}
p[right]{text-align: right;}
p[center]{text-align: center;}
```

```html
<p>простой абзац</p>
<p left>лево</p>
<p right>право</p>
<p center>центр</p>
<hr/>
<hr>
<b>жирный</b>
<i>наклонный</i>
<u>подчеркнутый</u>
<s>зачеркнутый</s>
<q>цитата</q>
<sub>в низ мелкий текст</sub>
<sup>в верх мелкий текст</sup>
```
