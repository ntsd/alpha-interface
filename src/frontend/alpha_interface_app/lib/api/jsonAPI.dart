import 'dart:convert';

import 'package:alpha_interface_app/model/yield.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/services.dart';

class JsonAPI {
  // creating a Singleton for the Api
  static final JsonAPI _singleton = new JsonAPI._internal();
  JsonAPI._internal();
  static JsonAPI getInstance() => _singleton;

  loadJson() async {
    List<Yield> yieldList = [];
    String data = await rootBundle.loadString("assets/rawData.json");
    List<dynamic> jsonResult = jsonDecode(data);

    jsonResult.forEach((v) { 
    Yield yi = new Yield.fromJson(v);
      yieldList.add(yi);
    });
    return yieldList;
  }
}
