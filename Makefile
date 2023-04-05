r:
	DIR=$(DIR) FILE=$(FILE) ./shell/run.sh

d:
	rm -rf ./test/ && URL=$(URL) ./shell/download.sh

t:
	DIR=$(DIR) FILE=$(FILE) ./shell/test.sh

s:
	DIR=$(DIR) FILE=$(FILE) URL=$(URL) ./shell/submit.sh

login:
	oj login https://atcoder.jp