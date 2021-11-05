class SmratContratAPI {
  // creating a Singleton for the Api
  static final SmratContratAPI _singleton = new SmratContratAPI._internal();
  SmratContratAPI._internal();
  static SmratContratAPI getInstance() => _singleton;

  initialize() async {
    //TODO: can be deleted if not needed
  }
}
