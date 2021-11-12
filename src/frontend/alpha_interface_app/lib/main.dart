import 'package:alpha_interface_app/screens/home.dart';
import 'package:alpha_interface_app/screens/login.dart';
import 'package:flutter/material.dart';

void main() async {
  WidgetsFlutterBinding
      .ensureInitialized(); // Make sure async call being awaited

  runApp(
    MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        backgroundColor: Colors.grey[600],
        appBarTheme: AppBarTheme(
          backgroundColor: Colors.grey[700],
        ),
        drawerTheme: DrawerThemeData(
          backgroundColor: Colors.grey[700],
        ),
      ),
      //home: Login(),
      home: HomePage("Alpha Interface Demo", "Haled"),
    ),
  );
}
