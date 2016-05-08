ID3 Algorithm in GoLang
-------
**Usage**
```
`go build` && ./id3
```
- don't forget to change file path for different file.

**file:** id3.go  **line:**  9
> dataset := Load("**data/baseball.csv**")

**Example Console Output**
```
	||-Outlook  =>  Overcast
	|			|->> Yes
	||-Outlook  =>  Sunny
	|			|-Humidity  =>  Normal
	|						|->> Yes
	|			|-Humidity  =>  High
	|						|->> No
	||-Outlook  =>  Rain
	|			|-Wind  =>  Weak
	|						|->> Yes
	|			|-Wind  =>  Strong
	|						|->> No
```


----------

 [The ID3 Algorithm Details](http://www.cise.ufl.edu/~ddd/cap6635/Fall-97/Short-papers/2.htm)