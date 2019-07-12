## Mini - A super minimalistic site generator

Sometimes a CMS or custom website seems to be overkill, welcome mini. A minimalistic site generator that aims to minimize configuration. All you need to do is initialize the directory of choice and you are good to go. For now it is up to you how you will serve the content. Currently you have to make your templates from scratch, mini will have 
default templates and template creation guidelines in the future. 

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

- [ ] General template themes
- [ ] Dev environment, edit and reload conveniently etc.
- [ ] Documentation
- [ ] Templates from dir (downloadable from GH)
- [x] Serve content "mini serve"
- [x] Cli for easy usage
