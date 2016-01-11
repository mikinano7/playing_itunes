iTunes = WScript.CreateObject("iTunes.Application");
 
WScript.ConnectObject(iTunes, "my_");
 
function my_OnPlayerPlayEvent(iTrack) {
    text = "NowPlaying: " + iTrack.Name + " - " + iTrack.Artist;
    WScript.Echo(text);
    return 0;
}
 
function my_OnQuittingEvent() {
    WScript.DisconnectObject(iTunes);
    WScript.Quit(0);
}
 
while (true) {
    WScript.sleep(1000);
}