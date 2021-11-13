class Order {
  int id;
  String traderName = "Haled";
  String traderLocation = "Berlin";
  String traderAdrres = "Example Street 10";
  String traderPostCode = "12345";
  String name;
  String tradeType;
  String actionType;
  int amount;
  int price;
  int duration;
  String createdAt;


  Order(this.id,this.name, this.tradeType, this.actionType, this.amount, this.price, this.createdAt, this.duration,{this.traderLocation});

}