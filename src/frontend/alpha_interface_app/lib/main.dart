
import 'package:alpha_interface_app/screens/home.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

void main() async {
  WidgetsFlutterBinding
      .ensureInitialized(); // Make sure async call being awaited

  runApp(
    MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData.dark(),
      home: HomePage('Alpha Interface Demo'),
      builder: EasyLoading.init(),
    ),
  );
}
