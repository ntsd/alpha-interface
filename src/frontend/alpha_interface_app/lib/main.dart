import 'package:alpha_interface_app/api/smart_contract_api.dart';
import 'package:alpha_interface_app/screens/home.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

// Global varaiable
SmratContratAPI _smratContratAPI = SmratContratAPI.getInstance();
void main() async {
  WidgetsFlutterBinding
      .ensureInitialized(); // Make sure async call being awaited

  await _smratContratAPI.initialize(); // initialize smartcontract singleton
  runApp(
    MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
          primaryColor: Colors.red,
          primarySwatch: Colors.red,
          accentColor: Colors.red),
      darkTheme: ThemeData.dark(),
      home: HomePage('Alpha Interface Demo'),
      builder: EasyLoading.init(),
    ),
  );
}
