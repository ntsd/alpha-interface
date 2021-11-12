import 'dart:math';

import 'package:alpha_interface_app/api/jsonAPI.dart';
import 'package:alpha_interface_app/model/order.dart';
import 'package:alpha_interface_app/model/yield.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:syncfusion_flutter_charts/charts.dart';

class FutureTrading extends StatefulWidget {
  final String productName;
  final String currentPrice;
  FutureTrading(this.productName, this.currentPrice);

  @override
  _FutureTradingState createState() => _FutureTradingState();
}

class _FutureTradingState extends State<FutureTrading> {
  JsonAPI _jsonAPI = JsonAPI.getInstance();
  List<Yield> yieldList = [];
  List<String> countries = [];
  String currentCountry = "";
  List<_SalesData> data = [];
  List<FuturePriceData> priceData = [];
  List<Order> orderList = [];
  int id = 0;
  int _amount = 100;
  int _price = 0;
  String location = "";
  int avaragePrice = 0;
  int totalPrice = 0;
  Random rand = Random();
  final TextEditingController _priceController = TextEditingController();

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    _priceController.text = "0";

    loadJson();
  }

  loadJson() async {
    yieldList = await _jsonAPI.loadCountryJson();
    currentCountry = yieldList.first.area;
    yieldList.forEach((element) {
      if (!countries.contains(element.area)) {
        countries.add(element.area);
      }
    });
    generateOrders();
    createChartDataForCountry();
    setState(() {});
  }

  generateOrders() {
    var actionType = "";

    if (currentCountry == "Germany") {
      location = "Berlin";
    } else {
      location = "Paris";
    }

    for (var i = 1; i < 10; i++) {
      id = i;
      int priceFlacuation = rand.nextInt(10);
      DateTime now = DateTime.now().add(Duration(days: -10 + i));
      String convertedDateTime =
          "${now.day.toString().padLeft(2, '0')}-${now.month.toString().padLeft(2, '0')}-${now.year.toString()} ${now.hour.toString().padLeft(2, '0')}-${now.minute.toString().padLeft(2, '0')}";
      int price = int.parse(widget.currentPrice) + priceFlacuation;

      bool randomBool = rand.nextBool();
      if (randomBool) {
        actionType = "Sell";
      } else {
        actionType = "Buy";
      }
      FuturePriceData newPriceData = FuturePriceData(convertedDateTime, price);
      priceData.add(newPriceData);
      totalPrice += price;
      int duration = rand.nextInt(3) + 1;
      Order newOrder = Order(id, widget.productName, "Future", actionType,
          _amount, price, convertedDateTime, duration,
          traderLocation: location);
      orderList.add(newOrder);
      orderList.sort((a, b) => a.createdAt.compareTo(b.createdAt));
    }
    avaragePrice = (totalPrice / orderList.length).round();
  }

  createChartDataForCountry() {
    data = [];
    List<_SalesData> newDataList = [];
    yieldList.forEach((element) {
      if (element.area == currentCountry) {
        print(element.area +
            " " +
            element.year.toString() +
            " : " +
            element.value.toString());
        _SalesData newData = _SalesData(element.year, element.value);
        newDataList.add(newData);
      }
    });

    data.addAll(newDataList.where((v) => data.every((d) => v.year != d.year)));
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
                                    createChartDataForCountry();
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
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      Container(
                        height: 400,
                        width: 800,
                        child: SfCartesianChart(
                            primaryXAxis: CategoryAxis(),
                            title: ChartTitle(
                                text: 'Yearly Production of ' +
                                    widget.productName),
                            tooltipBehavior:
                                TooltipBehavior(enable: true, header: ""),
                            series: <ChartSeries<_SalesData, String>>[
                              LineSeries<_SalesData, String>(
                                dataSource: data,
                                xValueMapper: (_SalesData sales, _) =>
                                    sales.year.toString(),
                                yValueMapper: (_SalesData sales, _) =>
                                    sales.value,
                              )
                            ]),
                      ),
                      SizedBox(
                        width: 50,
                      ),
                      Container(
                        height: 400,
                        width: 800,
                        child: SfCartesianChart(
                            primaryXAxis: CategoryAxis(),
                            title: ChartTitle(
                                text: 'Avarage Price $avaragePrice €'),
                            tooltipBehavior:
                                TooltipBehavior(enable: true, header: ""),
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
                    ],
                  ),
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
                          "Create Buy Order",
                          style: TextStyle(color: Colors.white),
                        ),
                        onPressed: () {
                          showOrderDialoge("Buy");
                        },
                      ),
                    ),
                    SizedBox(
                      width: 20,
                    ),
                    Container(
                      height: 50,
                      width: 300,
                      child: RaisedButton(
                        color: Colors.red,
                        child: Text(
                          "Create Sell Order",
                          style: TextStyle(color: Colors.white),
                        ),
                        onPressed: () {
                          showOrderDialoge("Sell");
                        },
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 10),
              ],
            ),
          ),
        ),
      ),
    );
  }

  showOrderDialoge(String action) {
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
                            initialValue: _amount.toString() + " kg",
                            enabled: false,
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
                          ),
                        ),
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
                          action, _amount, _price, convertedDateTime, duration,
                          traderLocation: location);
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
                  createTradeFields("Commodity: ", order.name),
                  createTradeFields("Type: ", order.tradeType),
                  createTradeFields(
                      "Duration: ", order.duration.toString() + " Years"),
                  createTradeFields(
                      "Amount: ", order.amount.toString() + " kg"),
                  createTradeFields("Price: ", order.price.toString() + " €"),
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
                  Text("Duration: "),
                  Text(order.duration.toString() + " Years"),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Type: "),
                  Text(order.tradeType),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Row(
                children: [
                  Text("Action: "),
                  Text(order.actionType),
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
                  Text("Postcode: "),
                  Text(order.traderPostCode),
                ],
              ),
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: RaisedButton(
                  color: Colors.red,
                  child: Text(
                    "Trade",
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
    countries.forEach((element) {
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
}

class _SalesData {
  _SalesData(this.year, this.value);

  final int year;
  final int value;
}

class FuturePriceData {
  FuturePriceData(this.time, this.price);

  final String time;
  final int price;
}
