import java.io.*;
import java.util.*;
import java.math.*;
import java.util.stream.*;

/**
 * Author :     AV
 * At   :       @Home
 * Date :       20.8.
 * ver. :       0.0
 * link :       
 * file :       cf/      
 */

// public class I_AM_AV{
class I_AM_AV{
    static private I_AM_AV_Template sc=new I_AM_AV_Template();
    static private StringBuilder sb=new StringBuilder();

    public static int code(int[] a, int n){
        return 0;
    }

    public static void main(String[] args) {
        int T=1;
        T=sc.nextInt();
        while(T-->0){
            // int n=sc.nextInt();
            // int[] a=sc.input(n);
            // int res = code(a, n);
            // sb.append(res + "\n");
        }
        System.out.print(sb);
    }
}

class I_AM_AV_Template{
    static private BufferedReader br;
    static private StringTokenizer st; 
    static public final int MOD=1000000007; //10^9+7

    I_AM_AV_Template(){
        br=new BufferedReader(new InputStreamReader(System.in));
    }  
    I_AM_AV_Template(String input)throws FileNotFoundException{
        System.setOut(new PrintStream(input));
    }
    I_AM_AV_Template(String input, String output)throws FileNotFoundException{
        System.setIn(new FileInputStream(input));
        System.setOut(new PrintStream(output));
    }
    /**
     * Use StringBuilder for output
     */
    String next() { 
        while (st == null || !st.hasMoreElements()) { 
            try{ 
                st = new StringTokenizer(br.readLine()); 
            } 
            catch (IOException e) { 
                e.printStackTrace(); 
            } 
        } 
        return st.nextToken(); 
    } 
    int nextInt() { return Integer.parseInt(next()); } 
    long nextLong() { return Long.parseLong(next()); } 
    double nextDouble() {  return Double.parseDouble(next()); } 
    String nextLine() { 
        String str = ""; 
        try{ 
            str = br.readLine(); 
        } 
        catch (IOException e) { 
            e.printStackTrace(); 
        } 
        return str; 
    }
    public int[] input(int n){
        int[] a=new int[n];
        for(int i=0;i<n;i++)
            a[i]=this.nextInt();
        return a;
    }
    public int sum(int[] a, int start, int end){
        return Arrays.stream(a, start, end).sum();
    }
    public <E>void display(E[] a, int start, int end){
        for (;start<end; start++) {
            System.out.print(a[start]+" ");
        }
        System.out.println();
    }
    public void displayMatrix(int[][] a){
        for(int[] x: a)
            System.out.println(Arrays.toString(x));
    }

    public int max(int[] arr){ return Arrays.stream(arr).max().getAsInt(); }
    public int min(int[] arr){ return Arrays.stream(arr).min().getAsInt(); }

    public <E> E max(E[] a, int n){
        E max=a[0];
        return max;
    }
    public boolean isPrime(int number){
        if(number <= 2)
            return number == 2;
        else
            return  (number % 2) != 0
                    &&
                IntStream.rangeClosed(3, (int) Math.sqrt(number))
                    .filter(n -> n % 2 != 0)
                    .noneMatch(n -> (number % n == 0));
        }
}


