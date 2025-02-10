import Toybox.Application;
import Toybox.Lang;
import Toybox.WatchUi;

class ExampleApp extends Application.AppBase {

    function initialize() {
        AppBase.initialize();
        var src = new ExampleMessage();
        src.i32 = 123;
        var byteArray = src.Encode();

        var dst = new ExampleMessage();
        dst.Decode(byteArray);
        System.print(dst.i32);
    }

    // onStart() is called on application start up
    function onStart(state as Dictionary?) as Void {
    }

    // onStop() is called when your application is exiting
    function onStop(state as Dictionary?) as Void {
    }

    // Return the initial view of your application here
    function getInitialView() as [Views] or [Views, InputDelegates] {
        return [ new WatchUi.View(), new WatchUi.BehaviorDelegate() ];
    }

}

function getApp() as ExampleApp {
    return Application.getApp() as ExampleApp;
}
