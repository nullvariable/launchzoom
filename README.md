just scratching my own itch here.

the install.sh file will do everything I needed to get this working on Ubuntu, YMMV.

## The problem
I created a custom browser launcher because I wanted to select which browser I used for which links (work slack into Chrome work profile, local tech slack into personal Chrome profile etc). One feature I added to this script was to just bypass all that with `xdg-open` for zoom links instead of click zoom link in slack, select browser, click ok to open in browser and then finally get a zoom link. Of course this created a new problem, anything that Slack (electron) opens ends up as a child of that process. I tinkered with a few different efforts, a nodejs script that tried to spawn it under a different PID and others, but never could get anything to work well. Interestingly, if I have FireFox or Chrome open, they seem to do a good job of coordinating and opening new tabs instead of child processes under Slack (not the case if they are closed and you open the from Slack, then Slack owns them 😢). So this was a good way to learn a bit about Go, and create a method where we can launch Zoom links really fast from Slack, but have them open in the main process instead of in a separate one (which fights with you and logs you out of your other instance)

## The solution
A user based systemd service that listens on a unix socket and launches `xdg-open` with `zoommtg://` links from the socket.
A cli app that takes any zoom url, base64 encodes it and writes it to that socket.

## Known issues
* This is a lot of stuff just to skip the browser step.
* Sometimes audio devices are missing from Zoom if the service launches Zoom (probably a permissions issue). I solved this by just ensuring that I launch Zoom and leave it running. There are likely other, perhaps better, options.

#### Useful links:
* https://vic.demuzere.be/articles/using-systemd-user-units/
* https://gobyexample.com/
* https://dave.dev/blog/2017/03/linux-systemd-golang-services-using-kardianos-service/
* https://www.howtogeek.com/216454/how-to-manage-systemd-services-on-a-linux-system/
* https://github.com/eliben/code-for-blog/blob/master/2019/unix-domain-sockets-go/simple-echo-server.go
* https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/
* https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/