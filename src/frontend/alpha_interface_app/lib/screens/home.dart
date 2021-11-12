import 'dart:math';

import 'package:alpha_interface_app/model/yield.dart';
import 'package:alpha_interface_app/screens/myassets.dart';
import 'package:alpha_interface_app/screens/futures.dart';
import 'package:alpha_interface_app/screens/tabBarScreen.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:syncfusion_flutter_charts/charts.dart';

class HomePage extends StatefulWidget {
  final String title;
  final String name;
  HomePage(this.title, this.name);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  List<Yield> yields = [];
  List<String> vegetables = [
    "Wheat",
    "Barely",
    "Maize",
    "Potatoes",
    "Tomatoes",
    "Strawberry",
    "Onions",
    "Banana"
  ];
  List<List<PriceData>> data = [];
  List<PriceData> totalData = [];
  List<int> firstPrices = [];
  List<int> lastPrices = [];
  int totalnewPrice = 0;

  var random = new Random();

  @override
  void initState() {
    super.initState();
    generateData();
    generateTotlatData();
  }

  generateData() {
    data = [];
    int generatedPrice = random.nextInt(100) + 25;

    vegetables.forEach((element) {
      List<PriceData> newDataList = [];
      for (var i = 1; i < 25; i++) {
        int price = generatedPrice + random.nextInt(20);
        int negetivePrice = random.nextInt(30);
        var result = price - negetivePrice;
        DateTime now = DateTime.now().add(Duration(hours: -25 + i));
        String convertedDateTime =
            "${now.hour.toString().padLeft(2, '0')}";

        PriceData newData = PriceData(result, convertedDateTime);
        newDataList.add(newData);
      }
      firstPrices.add(newDataList.first.value);
      print("This is first ${newDataList.first.value}");
      lastPrices.add(newDataList.last.value);
      print("This is last ${newDataList.last.value}");
      totalnewPrice += newDataList.last.value;
      data.add(newDataList);
    });

    return;
  }

  generateTotlatData() {
    List<Map<int, String>> map = [];
    totalData = [];
    data.forEach((element) {
      Map<int, String> newMap = {};
      element.forEach((sale) {
        newMap.addAll({sale.value: sale.time});
      });
      map.add(newMap);
    });
    for (var i = 1; i < 25; i++) {
      var total = 0;
      map.forEach((map) {
        map.forEach((key, value) {
          if (value == i.toString()) {
            total += key;
          }
        });
      });
      PriceData newData = PriceData(total, i.toString());
      totalData.add(newData);
    }
  }

  @override
  Widget build(BuildContext context) {
    var sidePadding = MediaQuery.of(context).size.width / 6;
    return Scaffold(
      appBar: AppBar(),
      body: SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.fromLTRB(sidePadding, 50, sidePadding, 50),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Padding(
                padding: const EdgeInsets.fromLTRB(35, 25, 8, 8),
                child: Text(
                  "Commodities (24h)",
                  style: TextStyle(
                      color: Colors.black,
                      fontSize: 24,
                      fontWeight: FontWeight.bold),
                ),
              ),
              Padding(
                padding: const EdgeInsets.all(35.0),
                child: Container(
                  child: GridView.builder(
                    itemCount: vegetables.length,
                    gridDelegate:
                        const SliverGridDelegateWithFixedCrossAxisCount(
                      crossAxisCount: 4,
                    ),
                    shrinkWrap: true,
                    itemBuilder: (context, index) {
                      var priceDiff;
                      var lastPrice = lastPrices[index];
                      var firstPrice = firstPrices[index];
                      priceDiff = lastPrice - firstPrice;
                      return GestureDetector(
                        onTap: () {
                          Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (BuildContext cntx) => MyTabBar(
                                      vegetables[index],
                                      lastPrice.toString(),
                                      data[index])));
                        },
                        child: Card(
                          shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(10.0)),
                          elevation: 2,
                          child: Container(
                            height: 100,
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Padding(
                                  padding:
                                      const EdgeInsets.fromLTRB(20, 20, 20, 10),
                                  child: Text(
                                    vegetables[index],
                                    style: TextStyle(
                                      color: Colors.black,
                                      fontSize: 24,
                                    ),
                                  ),
                                ),
                                Flexible(
                                  flex: 3,
                                  child: SfCartesianChart(
                                      primaryXAxis: CategoryAxis(
                                          title: AxisTitle(
                                              text: 'Time',
                                              textStyle:
                                                  TextStyle(fontSize: 12))),
                                      primaryYAxis: NumericAxis(
                                          title: AxisTitle(
                                              text: 'Price',
                                              textStyle:
                                                  TextStyle(fontSize: 12))),
                                      tooltipBehavior:
                                          TooltipBehavior(enable: false),
                                      series: <ChartSeries<PriceData, dynamic>>[
                                        LineSeries<PriceData, String>(
                                          dataSource: data[index],
                                          xValueMapper: (PriceData sales, _) =>
                                              sales.time,
                                          yValueMapper: (PriceData sales, _) =>
                                              sales.value,
                                        )
                                      ]),
                                ),
                                Padding(
                                  padding: const EdgeInsets.only(left: 20),
                                  child: Text(
                                    lastPrice.toString() + " €",
                                    style: TextStyle(
                                        color: Colors.black,
                                        fontSize: 24,
                                        fontWeight: FontWeight.bold),
                                  ),
                                ),
                                Flexible(
                                  flex: 1,
                                  child: Padding(
                                    padding: const EdgeInsets.fromLTRB(
                                        20, 10, 20, 0),
                                    child: Row(
                                      children: [
                                        Text(
                                          priceDiff >= 0
                                              ? "+ $priceDiff €"
                                              : "$priceDiff €",
                                          style: TextStyle(
                                            color: priceDiff >= 0
                                                ? Colors.green
                                                : Colors.red,
                                            fontSize: 12,
                                          ),
                                        ),
                                        Icon(
                                          priceDiff >= 0
                                              ? Icons.keyboard_arrow_up
                                              : Icons.keyboard_arrow_down,
                                          color: priceDiff >= 0
                                              ? Colors.green
                                              : Colors.red,
                                        ),
                                      ],
                                    ),
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ),
                      );
                    },
                  ),
                ),
              ),
              Padding(
                padding: const EdgeInsets.only(top: 10, bottom: 10),
                child: Container(
                  width: MediaQuery.of(context).size.width,
                  height: 1,
                  color: Colors.grey[300],
                ),
              ),
              //buildMyPortfolio()
            ],
          ),
        ),
      ),
    );
  }

  buildMyPortfolio() {
    return Container(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Padding(
            padding: const EdgeInsets.fromLTRB(35, 35, 20, 10),
            child: Text(
              "My Portfolio",
              style: TextStyle(
                color: Colors.black,
                fontSize: 24,
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.fromLTRB(35, 10, 35, 10),
            child: Text(
              totalnewPrice.toString() + " €",
              style: TextStyle(
                  color: Colors.black,
                  fontSize: 35,
                  fontWeight: FontWeight.bold),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(right: 35.0, left: 20),
            child: Container(
              height: 300,
              child: SfCartesianChart(
                  primaryXAxis: CategoryAxis(
                      title: AxisTitle(
                          text: 'Time', textStyle: TextStyle(fontSize: 12))),
                  primaryYAxis: NumericAxis(
                      title: AxisTitle(
                          text: 'Price', textStyle: TextStyle(fontSize: 12))),
                  tooltipBehavior: TooltipBehavior(enable: false),
                  series: <ChartSeries<PriceData, dynamic>>[
                    LineSeries<PriceData, String>(
                      dataSource: totalData,
                      xValueMapper: (PriceData sales, _) => sales.time,
                      yValueMapper: (PriceData sales, _) => sales.value,
                    )
                  ]),
            ),
          ),
        ],
      ),
    );
  }
}

class PriceData {
  PriceData(
    this.value,
    this.time,
  );

  final String time;
  final int value;
}
