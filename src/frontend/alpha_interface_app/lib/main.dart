
import 'package:alpha_interface_app/screens/home.dart';
import 'package:flutter/material.dart';

void main() async {
  WidgetsFlutterBinding
      .ensureInitialized(); // Make sure async call being awaited

  runApp(
    MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData.dark(),
      home: HomePage('Alpha Interface Demo'),
    ),
  );
}
