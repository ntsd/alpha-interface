import 'package:alpha_interface_app/screens/home.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class Login extends StatefulWidget {
  final String productName;
  Login(this.productName);

  @override
  _LoginState createState() => _LoginState();
}

class _LoginState extends State<Login> {
  String type = "";
  final TextEditingController _typeFieldController = TextEditingController();
  @override
  void initState() {
    super.initState();
    _typeFieldController.text = "";
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[200],
      body: Center(
        child: Container(
          width: MediaQuery.of(context).size.width / 3,
          height: 300,
          child: Card(
            shape:
                RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
            color: Colors.grey[700],
            elevation: 10,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                SizedBox(
                  height: 25,
                ),
                Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: Text(
                    "Name",
                    style: TextStyle(color: Colors.white, fontSize: 24),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: Container(
                    child: TextFormField(
                      style: TextStyle(color: Colors.white),
                      cursorColor: Colors.red,
                      decoration: InputDecoration(
                          focusedBorder: OutlineInputBorder(
                            borderSide:
                                BorderSide(color: Colors.red, width: 2.0),
                          ),
                          enabledBorder: OutlineInputBorder(
                            borderSide:
                                BorderSide(color: Colors.red, width: 2.0),
                          ),
                          hintText: 'Enter your name',
                          hintStyle: TextStyle(color: Colors.white)),
                      controller: _typeFieldController,
                      onChanged: (value) {
                        setState(() {
                          type = value;
                        });
                      },
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: Container(
                    height: 50,
                    width: MediaQuery.of(context).size.width / 3,
                    child: RaisedButton(
                      color: Colors.red,
                      child: Text(
                        "Login",
                        style: TextStyle(color: Colors.white),
                      ),
                      onPressed: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (BuildContext cntx) =>
                                HomePage(widget.productName, type),
                          ),
                        );
                      },
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
