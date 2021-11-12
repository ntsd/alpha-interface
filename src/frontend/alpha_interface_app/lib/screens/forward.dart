import 'dart:math';

import 'package:alpha_interface_app/api/jsonAPI.dart';
import 'package:alpha_interface_app/model/order.dart';
import 'package:alpha_interface_app/model/yield.dart';
import 'package:alpha_interface_app/screens/futures.dart';
import 'package:alpha_interface_app/screens/home.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:syncfusion_flutter_charts/charts.dart';

class ForwardTrading extends StatefulWidget {
  final String productName;
  final String currentPrice;
  ForwardTrading(this.productName, this.currentPrice);

  @override
  _ForwardTradingState createState() => _ForwardTradingState();
}

class _ForwardTradingState extends State<ForwardTrading> {
  JsonAPI _jsonAPI = JsonAPI.getInstance();
  List<Yield> yieldList = [];
  String currentCountry = "";
  List<String> cities = [];
  List<Order> orderList = [];
  List<FuturePriceData> priceData = [];
  int id = 0;
  int _amount = 10;
  int _price = 0;
  int avaragePrice = 0;
  int totalPrice = 0;
  Random rand = Random();
  final TextEditingController _priceController = TextEditingController();
  int buyAmount = 0;
  int sellAmount = 0;

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    loadJson();
    _priceController.text = "0";
  }

  loadJson() async {
    yieldList = await _jsonAPI.loadJson();
    currentCountry = yieldList.first.area;
    yieldList.forEach((element) {
      if (!cities.contains(element.area)) {
        cities.add(element.area);
      }
    });
    generateOrders();
    setState(() {});
  }

  generateOrders() {
    var actionType = "";

    for (var i = 1; i < 10; i++) {
      id = i;
      int priceFlacuation = rand.nextInt(2);
      DateTime now = DateTime.now().add(Duration(days: -10 + i));
      String convertedDateTime =
          "${now.day.toString().padLeft(2, '0')}-${now.month.toString().padLeft(2, '0')}-${now.year.toString()} ${now.hour.toString().padLeft(2, '0')}-${now.minute.toString().padLeft(2, '0')}";
      int price =
          (int.parse(widget.currentPrice) / 10).round() + priceFlacuation;

      FuturePriceData newPriceData = FuturePriceData(convertedDateTime, price);
      priceData.add(newPriceData);
      totalPrice += price;
      int duration = rand.nextInt(3);
      Order newOrder = Order(id, widget.productName, "Future", actionType,
          _amount, price, convertedDateTime, duration,
          traderLocation: currentCountry);
      orderList.add(newOrder);
      orderList.sort((a, b) => a.createdAt.compareTo(b.createdAt));
    }
    avaragePrice = (totalPrice / orderList.length).round();
  }

  @override
  Widget build(BuildContext context) {
    var sidePadding = MediaQuery.of(context).size.width / 6;
    return Scaffold(
      body: SingleChildScrollView(
        child: Container(
          width: MediaQuery.of(context).size.width,
          child: Padding(
            padding: EdgeInsets.fromLTRB(sidePadding, 50, sidePadding, 50),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    Flexible(
                      flex: 2,
                      child: Padding(
                        padding: const EdgeInsets.all(10.0),
                        child: Container(
                          height: 50,
                          decoration: BoxDecoration(
                            color: Colors.white,
                            border: Border.all(width: 1),
                            borderRadius: BorderRadius.circular(2),
                          ),
                          child: Center(
                              child: Text(
                            widget.productName,
                            style: TextStyle(fontSize: 30, color: Colors.black),
                          )),
                        ),
                      ),
                    ),
                    Flexible(
                      flex: 3,
                      child: Padding(
                        padding: const EdgeInsets.all(10.0),
                        child: Container(
                          decoration: BoxDecoration(
                            color: Colors.white,
                            border: Border.all(width: 1),
                            borderRadius: BorderRadius.circular(2),
                          ),
                          child: DropdownButtonHideUnderline(
                            child: ButtonTheme(
                              alignedDropdown: true,
                              child: DropdownButton<String>(
                                iconEnabledColor: Colors.black,
                                value: currentCountry,
                                isExpanded: true,
                                underline: Container(),
                                onChanged: (String value) {
                                  setState(() {
                                    currentCountry = value;
                                    orderList = [];
                                    priceData = [];
                                    avaragePrice = 0;
                                    totalPrice = 0;
                                    generateOrders();
                                  });
                                },
                                items: createDropDownItems(),
                              ),
                            ),
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
                Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: SfCartesianChart(
                      primaryXAxis: CategoryAxis(),
                      title: ChartTitle(text: 'Avarage Price $avaragePrice €'),
                      tooltipBehavior: TooltipBehavior(enable: true),
                      series: <ChartSeries<FuturePriceData, String>>[
                        LineSeries<FuturePriceData, String>(
                          dataSource: priceData,
                          xValueMapper: (FuturePriceData sales, _) =>
                              sales.time,
                          yValueMapper: (FuturePriceData sales, _) =>
                              sales.price,
                        )
                      ]),
                ),
                Padding(
                  padding: const EdgeInsets.only(top: 25.0),
                  child: Container(
                    decoration: BoxDecoration(
                        border: Border(
                            top: BorderSide(width: 2, color: Colors.red),
                            right: BorderSide(width: 2, color: Colors.red),
                            left: BorderSide(width: 2, color: Colors.red))),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: [
                        ListView.builder(
                          shrinkWrap: true,
                          itemCount: orderList.length,
                          itemBuilder: (context, index) {
                            return createOrderWidget(orderList[index], index);
                          },
                        ),
                      ],
                    ),
                  ),
                ),
                SizedBox(
                  height: 25,
                ),
                Row(
                  children: [
                    Container(
                      height: 50,
                      width: 300,
                      child: RaisedButton(
                        color: Colors.red,
                        child: Text(
                          "Create Order",
                          style: TextStyle(color: Colors.white),
                        ),
                        onPressed: () {
                          showOrderDialoge();
                        },
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  showTradeDialoge(Order order, int index) {
    showDialog(
        context: context,
        builder: (context) {
          return AlertDialog(
            title: Text("Trade"),
            content: Container(
              height: 600,
              width: 400,
              child: Column(
                children: [
                  createTradeFields("Creator: ", order.traderName),
                  createTradeFields("Location: ", order.traderLocation),
                  createTradeFields("Address: ", order.traderAdrres),
                  createTradeFields("Postcode: ", order.traderPostCode),
                  createTradeFields("Commodity: ", order.name),
                  createTradeFields("Amount: ", order.amount.toString()),
                  createTradeFields("Creator: ", order.price.toString()),
                  createTradeFields("Created at: ", order.createdAt.toString()),
                ],
              ),
            ),
            actions: [
              Padding(
                padding: const EdgeInsets.all(10.0),
                child: Container(
                  height: 50,
                  child: RaisedButton(
                    color: Colors.grey[700],
                    child: Text(
                      "Cancel",
                      style: TextStyle(color: Colors.white),
                    ),
                    onPressed: () {
                      Navigator.pop(context);
                    },
                  ),
                ),
              ),
              Padding(
                padding: const EdgeInsets.all(10.0),
                child: Container(
                  height: 50,
                  child: RaisedButton(
                    color: Colors.red,
                    child: Text(
                      "Execute Trade",
                      style: TextStyle(color: Colors.white),
                    ),
                    onPressed: () {
                      orderList.removeAt(index);
                      Navigator.pop(context);
                      setState(() {});
                    },
                  ),
                ),
              ),
            ],
          );
        });
  }

  createTradeFields(String key, String value) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Row(
        children: [
          Container(width: 100, child: Text("$key: ")),
          SizedBox(
            width: 5,
          ),
          Flexible(
            child: TextFormField(
              initialValue: value,
              enabled: false,
              style: TextStyle(color: Colors.red),
              decoration: InputDecoration(
                focusedBorder: OutlineInputBorder(
                  borderSide: BorderSide(color: Colors.red, width: 2.0),
                ),
                enabledBorder: OutlineInputBorder(
                  borderSide: BorderSide(color: Colors.red, width: 2.0),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  createOrderWidget(Order order, int index) {
    return Column(
      children: [
        Padding(
          padding: const EdgeInsets.only(left: 8.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Row(
                children: [
                  Text("Commodity: "),
                  Text(order.name),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Amount: "),
                  Text(order.amount.toString() + " kg"),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Price: "),
                  Text(order.price.toString() + " €"),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Created at: "),
                  Text(order.createdAt.toString()),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Location: "),
                  Text(order.traderLocation),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Address: "),
                  Text(order.traderAdrres),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Postcode: "),
                  Text(order.traderPostCode),
                ],
              ),
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: RaisedButton(
                  color: Colors.red,
                  child: Text(
                    "Buy",
                    style: TextStyle(color: Colors.white),
                  ),
                  onPressed: () {
                    showTradeDialoge(order, index);
                  },
                ),
              ),
            ],
          ),
        ),
        SizedBox(
          height: 5,
        ),
        Container(
          height: 2,
          color: Colors.red,
        ),
      ],
    );
  }

  createDropDownItems() {
    List<DropdownMenuItem<String>> list = [];
    cities.forEach((element) {
      var item = DropdownMenuItem<String>(
          value: element,
          child: Container(
            child: Center(
                child: Text(
              element,
              style: TextStyle(fontSize: 30, color: Colors.black),
            )),
          ));
      list.add(item);
    });
    return list;
  }

  showOrderDialoge() {
    showDialog(
        context: context,
        builder: (context) {
          return AlertDialog(
            title: Text("Create Order"),
            content: Container(
              height: 150,
              width: 300,
              child: Column(
                children: [
                  Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Row(
                      children: [
                        Container(width: 100, child: Text("Amount: ")),
                        SizedBox(
                          width: 5,
                        ),
                        Flexible(
                            child: TextFormField(
                              enabled: false,
                          initialValue: _amount.toString(),
                          style: TextStyle(color: Colors.red),
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
                          ),
                        )),
                      ],
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Row(
                      children: [
                        Container(width: 100, child: Text("Price: ")),
                        SizedBox(
                          width: 5,
                        ),
                        Flexible(
                          child: TextFormField(
                            style: TextStyle(color: Colors.red),
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
                                hintText: 'Enter price',
                                hintStyle: TextStyle(color: Colors.red)),
                            controller: _priceController,
                            onChanged: (value) {
                              setState(() {
                                _price = int.parse(value);
                              });
                            },
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
            actions: [
              Padding(
                padding: const EdgeInsets.all(10.0),
                child: Container(
                  height: 50,
                  child: RaisedButton(
                    color: Colors.red,
                    child: Text(
                      "Create Order",
                      style: TextStyle(color: Colors.white),
                    ),
                    onPressed: () {
                      id += 1;
                      DateTime now =
                          DateTime.now().add(Duration(days: -10 + id));
                      String convertedDateTime =
                          "${now.day.toString().padLeft(2, '0')}-${now.month.toString().padLeft(2, '0')}-${now.year.toString()} ${now.hour.toString().padLeft(2, '0')}-${now.minute.toString().padLeft(2, '0')}";
                      FuturePriceData newPriceData =
                          FuturePriceData(convertedDateTime, _price);
                      priceData.add(newPriceData);
                      totalPrice += _price;
                      avaragePrice = 0;
                      avaragePrice = (totalPrice / orderList.length).round();
                      int duration = rand.nextInt(3);
                      Order newOrder = Order(id, widget.productName, "Future",
                          "Sell", _amount, _price, convertedDateTime, duration,
                          traderLocation: currentCountry);
                      orderList.add(newOrder);
                      Navigator.pop(context);
                      setState(() {});
                    },
                  ),
                ),
              ),
            ],
          );
        });
  }
}

class _SalesData {
  _SalesData(this.year, this.value);

  final int year;
  final int value;
}
