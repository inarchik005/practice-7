public class TestTime {
    public static void main(String[] args) {

        Time t1 = new Time(23, 5, 6);
        Time t2 = new Time(4, 24, 33);

        t1.add(t2);

        System.out.println("Universal: " + t1.toUniversal());
        System.out.println("Standard: " + t1.toStandard());
    }
}
