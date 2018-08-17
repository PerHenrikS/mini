## Mini - A super minimalistic static site generator

Sometimes a CMS or custom website seems to be overkill, welcome mini. A minimalistic static site generator (at the moment more like a blog generator) that aims to minimize config. All you need to do is initialize the directory of choice and you are good to go. For now it is up to you how you will serve the content.

## Usage 

Dependencies: 
```
Go
git 
```

Download and build: 
```
cd into your $GOPATH/src 
git clone https://github.com/PerHenrikS/mini
cd mini
go install
```

Initialize and generate: 
```
mkdir mywebsite
cd mywebsite/ 
mini init
mini gen
mini serve
```


### Planned features

- [ ] Template inheritance for more complicated layouts
- [ ] Easier template themes
- [ ] Different templates (pages, posts, something...)
- [ ] Dev environment, edit and reload conveniently etc.
- [ ] Documentation
- [x] Serve content "mini serve"
- [x] Cli for ez usage
- [x] Command: init - initialize directory 
