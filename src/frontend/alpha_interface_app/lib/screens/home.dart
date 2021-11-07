
import 'package:alpha_interface_app/model/yield.dart';
import 'package:alpha_interface_app/screens/product_info.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class HomePage extends StatefulWidget {
  final String title;
  HomePage(this.title);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  List<Yield> yields = [];
  List<String> vegetables = ["Wheat", "Barely", "Maize", "Potatoes"];

  @override
  void initState() {
    super.initState();
  }



  @override
  Widget build(BuildContext context) {
    var widthOneThird = MediaQuery.of(context).size.width / 3;
    return Scaffold(
      appBar: AppBar(
        title: Center(
          child: Text(widget.title),
        ),
      ),
      body: ListView.builder(
        itemCount: vegetables.length,
        shrinkWrap: true,
        itemBuilder: (context, index) {
          return GestureDetector(
            onTap: () {
              Navigator.push(
                  context,
                  MaterialPageRoute(
                      builder: (BuildContext cntx) =>
                          ProductInfo(vegetables[index])));
            },
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: Container(
                height: 100,
                decoration: BoxDecoration(
                  color: Colors.white,
                  border: Border.all(width: 1),
                  borderRadius: BorderRadius.circular(2),
                ),
                child: Center(
                  child: Container(
                    child: Text(
                      vegetables[index],
                      style: TextStyle(
                          color: Colors.black,
                          fontSize: 24,
                          fontWeight: FontWeight.bold),
                    ),
                  ),
                ),
              ),
            ),
          );
        },
      ),
    );
  }
}
