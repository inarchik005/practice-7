public class Student {
    private String name;
    private String id;
    private int year;

    public Student(String name, String id, int year) {
        this.name = name;
        this.id = id;
        this.year = year;
    }

    public String getName() { return name; }
    public String getId() { return id; }
    public int getYear() { return year; }

    public void nextYear() { year++; }

    public static void main(String[] args) {
        Student s = new Student("Inara", "23b031224", 3);

        System.out.println(s.getName());
        System.out.println(s.getId());
        System.out.println(s.getYear());

        s.nextYear();
        System.out.println("After next year:");
        System.out.println(s.getYear());
    }
}
